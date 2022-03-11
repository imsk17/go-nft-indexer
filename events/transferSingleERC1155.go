package events

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransferSingleERC1155 represents a single transfer event in any ERC1155 smart contract. Emitted when `tokenId` token is transferred from `from` to `to`.
type TransferSingleERC1155 struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Contract string
}

// Topic returns the event's topic (keckakk256('EventSignature')).
func (e TransferSingleERC1155) Topic() string {
	return "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
}

// String returns a human readable string representation of the event.
func (e *TransferSingleERC1155) String() string {
	return fmt.Sprintf("Transfer Single :- Operator : %s, Value: %v, ID: %v\n%s -> %s\n", e.Operator, e.Value.Int64(), e.Id.Int64(), e.From, e.To)
}

// This function can be used to decode a Transfer event from an event log.
func DecodeTransferSingleERC1155(log *types.Log) TransferSingleERC1155 {
	var transferSingle TransferSingleERC1155
	err := ERC1155ABI.UnpackIntoInterface(&transferSingle, "TransferSingle", log.Data)
	if err != nil {
		fmt.Printf("Error Parsing TransferSingle : %e\n", err)
	}
	transferSingle.Operator = common.BytesToAddress(log.Topics[1].Bytes())
	transferSingle.From = common.BytesToAddress(log.Topics[2].Bytes())
	transferSingle.To = common.BytesToAddress(log.Topics[3].Bytes())
	transferSingle.Contract = log.Address.Hex()
	return transferSingle
}
