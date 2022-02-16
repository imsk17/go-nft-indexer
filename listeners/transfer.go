package listeners

import (
	"github.com/ethereum/go-ethereum/core/types"
	"go-nft-listener/events"
)

type transferListener struct {
	ChainName    string
	Rpc          string
	EventChannel chan<- events.Event
	LogChannel   <-chan types.Log
}

func NewTransfer(chainName string, rpc string, eventChannel chan<- events.Event, logChannel <-chan types.Log) Listener {
	return &transferListener{chainName, rpc, eventChannel, logChannel}
}

func (l *transferListener) Listen() {
	for log := range l.LogChannel {
		// Is a TransferERC721/TransferSingleERC1155/TransferBatchERC1155 Event
		if len(log.Topics) > 0 {
			if log.Topics[0].String() == (events.TransferERC721{}).Topic() && len(log.Topics) == 4 {
				ev := events.DecodeTransferERC721(&log)
				l.EventChannel <- &ev
			}

			if log.Topics[0].String() == (events.TransferSingleERC1155{}).Topic() && len(log.Topics) == 4 {
				ev := events.DecodeTransferSingleERC1155(&log)
				l.EventChannel <- &ev
			}

			if log.Topics[0].String() == (events.TransferBatchERC1155{}).Topic() && len(log.Topics) == 4 {
				ev := events.DecodeTransferBatchERC1155(&log)
				l.EventChannel <- &ev
			}
		}
	}
}

func (l *transferListener) Close() {
	close(l.EventChannel)
}
