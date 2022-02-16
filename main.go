package main

import (
	"context"
	"fmt"
	"go-nft-listener/config"
	"go-nft-listener/events"
	"go-nft-listener/listeners"
	"sync"

	log "github.com/sirupsen/logrus"

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

	log.SetLevel(c.LogLevel)

	log.Info("☎️ Connecting to Chain ...")
	client, err := ethclient.Dial(c.Rpc)
	if err != nil {
		log.Panic(err)
	}

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
		fmt.Println(<-eventChan)
	}
}
