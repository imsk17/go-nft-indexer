package main

import (
	"context"
	"go-nft-listener/config"
	"go-nft-listener/domain"
	"go-nft-listener/events"
	"go-nft-listener/listeners"
	"sync"

	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var wg *sync.WaitGroup

func main() {
	log.Info("🚀 Starting NFT Listener")

	// Load Environment Variables
	log.Info("⚙️ Loading Environment Variables")
	c := config.Load()

	log.Info("✅ Environment Variables Loaded")

	log.SetLevel(c.LogLevel)

	// Create Connection to Database
	log.Info("💾 Connecting to Database ...")
	db, err := gorm.Open(postgres.Open(c.Db), &gorm.Config{
		Logger: gorm_logrus.New(),
	})
	if err != nil {
		log.Panicf("💥 Failed to connect to database: %s", err)
	}
	log.Info("✅ Connected to Database")

	if c.LogLevel == log.DebugLevel {
		log.Info("💾 Switching ORM To Debug Mode")
		db.Debug()
		log.Info("✅ Database Debug Mode On")
	}

	log.Info("💾 Trying to auto migrate the database ...")
	err = db.AutoMigrate(&domain.EthNft{})

	if err != nil {
		log.Panicf("💥 Failed to auto migrate the database: %s", err)
	}

	log.Info("☎️ Connecting to Chain ...")

	client, err := ethclient.Dial(c.Rpc)
	if err != nil {
		log.Panic("💥 Failed to connect to chain: %s", err)
	}
	log.Info("✅ Connected to Chain")

	chainId, err := client.ChainID(context.Background())

	if err != nil {
		log.Panic("💥 Failed to get chain id: %s", err)
	}

	log.Infof("✅ Connection to Chain Successful. Chain ID: %v", chainId)

	logChan := make(chan types.Log)

	if _, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{
				common.HexToHash(events.TransferERC721{}.Topic()),
				common.HexToHash(events.TransferSingleERC1155{}.Topic()),
				common.HexToHash(events.TransferBatchERC1155{}.Topic()),
			},
		},
	}, logChan); err != nil {
		log.Panicf("💥 Failed to subscribe to logs: %s", err)
	}

	eventChan := make(chan events.Event)

	listener := listeners.NewTransfer("Polygon", c.Rpc, eventChan, logChan, client)

	go func() {
		listener.Listen()
		wg.Add(1)
	}()

	for {
		<-eventChan
	}
}
