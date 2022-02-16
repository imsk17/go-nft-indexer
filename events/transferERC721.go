package events

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TransferERC721 struct {
	From     common.Address
	To       common.Address
	TokenId  *big.Int
	Contract string
}

func (e TransferERC721) Topic() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}

func (e *TransferERC721) String() string {
	return fmt.Sprintf("Transfer ERC721 :-\n%s -> %s, Token: %v\n", e.From.Hex(), e.To.Hex(), e.TokenId)
}

func DecodeTransferERC721(log *types.Log) TransferERC721 {
	var transfer TransferERC721
	transfer.From = common.BytesToAddress(log.Topics[1].Bytes())
	transfer.To = common.BytesToAddress(log.Topics[2].Bytes())
	transfer.TokenId = log.Topics[3].Big()
	transfer.Contract = log.Address.Hex()
	return transfer
}
