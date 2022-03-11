package listeners

import (
	"go-nft-listener/events"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc165"
	log "github.com/sirupsen/logrus"
)

type transferListener struct {
	ChainName    string
	Rpc          string
	EventChannel chan<- events.Event
	LogChannel   <-chan types.Log
	Client       *ethclient.Client
}

func NewTransfer(chainName string, rpc string, eventChannel chan<- events.Event, logChannel <-chan types.Log, c *ethclient.Client) Listener {
	return &transferListener{chainName, rpc, eventChannel, logChannel, c}
}

func (l *transferListener) Listen() {
	for eventLog := range l.LogChannel {
		// Is a TransferERC721/TransferSingleERC1155/TransferBatchERC1155 Event
		if len(eventLog.Topics) > 0 {
			if eventLog.Topics[0].String() == (events.TransferERC721{}).Topic() && len(eventLog.Topics) == 4 {
				erc16, err := erc165.NewErc165(eventLog.Address, l.Client)
				if err != nil {
					log.Warnf("Failed to connect to ERC165 Interface: %s", err.Error())
				}
				supports, err := erc16.SupportsInterface(nil, erc165.InterfaceIdErc721)
				if err != nil {
					log.Warnf("Failed to query if the contract is ERC165 supported or not: %s", err.Error())
				}
				if supports {
					ev := events.DecodeTransferERC721(&eventLog)
					l.EventChannel <- &ev
				}
			}

			if eventLog.Topics[0].String() == (events.TransferSingleERC1155{}).Topic() && len(eventLog.Topics) == 4 {
				erc16, err := erc165.NewErc165(eventLog.Address, l.Client)
				if err != nil {
					log.Warnf("Failed to connect to ERC165 Interface: %s", err.Error())
				}
				supports, err := erc16.SupportsInterface(nil, erc165.InterfaceIdErc1155)
				if err != nil {
					log.Warnf("Failed to query if the contract is ERC165 supported or not: %s", err.Error())
				}
				if supports {
					ev := events.DecodeTransferSingleERC1155(&eventLog)
					l.EventChannel <- &ev
				}
			}

			if eventLog.Topics[0].String() == (events.TransferBatchERC1155{}).Topic() && len(eventLog.Topics) == 4 {
				erc16, err := erc165.NewErc165(eventLog.Address, l.Client)
				if err != nil {
					log.Warnf("Failed to connect to ERC165 Interface: %s", err.Error())
				}
				supports, err := erc16.SupportsInterface(nil, erc165.InterfaceIdErc1155)
				if err != nil {
					log.Warnf("Failed to query if the contract is ERC165 supported or not: %s", err.Error())
				}
				if supports {
					ev := events.DecodeTransferBatchERC1155(&eventLog)
					l.EventChannel <- &ev
				}
			}
		}
	}
}

func (l *transferListener) Close() {
	close(l.EventChannel)
}
