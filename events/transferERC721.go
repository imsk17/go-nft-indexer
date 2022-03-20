package events

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransferERC721 represents a single transfer event in any ERC721 smart contract. Emitted when `tokenId` token is transferred from `from` to `to`.
type TransferERC721 struct {
	From     common.Address
	To       common.Address
	TokenId  *big.Int
	Contract string
}

// Topic returns the event's topic (keckakk256('EventSignature')).
func (e TransferERC721) Topic() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}

// String returns a human-readable string representation of the event.
func (e *TransferERC721) String() string {
	return fmt.Sprintf("Transfer ERC721 :-\n%s -> %s, Token: %v\n", e.From.Hex(), e.To.Hex(), e.TokenId)
}

// DecodeTransferERC721 can be used to decode a Transfer event from an event log.
func DecodeTransferERC721(log *types.Log) TransferERC721 {
	var transfer TransferERC721
	transfer.From = common.BytesToAddress(log.Topics[1].Bytes())
	transfer.To = common.BytesToAddress(log.Topics[2].Bytes())
	transfer.TokenId = log.Topics[3].Big()
	transfer.Contract = log.Address.Hex()
	return transfer
}
