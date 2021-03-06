package events

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/metachris/eth-go-bindings/erc1155"
)

var ERC1155ABI, _ = abi.JSON(strings.NewReader(erc1155.Erc1155ABI))

// TransferBatchERC1155 represents a batch transfer event in any ERC1155 smart contract. Equivalent to
// multiple TransferSingle events, where operator, from and to are the same for all transfers.
type TransferBatchERC1155 struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Contract string
}

// Topic returns the event's topic (keckakk256('EventSignature')).
func (e TransferBatchERC1155) Topic() string {
	return "0xf5f16c58bf69e14e9fa06e742215b42aa896de1c15af339f09e3360557089f43"
}

// String returns a human-readable string representation of the event.
func (e *TransferBatchERC1155) String() string {
	return fmt.Sprintf("Transfer ERC1155:\nOperator: %s\n%s -> %s\n Value: %v Ids: %v", e.Operator, e.From, e.To, e.Values, e.Ids)
}

// DecodeTransferBatchERC1155 can be used to decode a TransferBatchERC1155 event from an event log.
func DecodeTransferBatchERC1155(log *types.Log) TransferBatchERC1155 {
	var transferBatch TransferBatchERC1155
	err := ERC1155ABI.UnpackIntoInterface(&transferBatch, "TransferBatch", log.Data)
	if err != nil {
		fmt.Printf("Error Parsing TransferBatch : %e\n", err)
	}
	transferBatch.Operator = common.BytesToAddress(log.Topics[1].Bytes())
	transferBatch.From = common.BytesToAddress(log.Topics[2].Bytes())
	transferBatch.To = common.BytesToAddress(log.Topics[3].Bytes())
	transferBatch.Contract = log.Address.Hex()
	return transferBatch
}
