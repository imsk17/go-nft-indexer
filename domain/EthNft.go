package domain

import (
	"errors"

	"gorm.io/gorm"
)

type Contract string

const (
	ERC721  Contract = "ERC721"
	ERC1155 Contract = "ERC1155"
)

type EthNft struct {
	gorm.Model
	ChainId      string   `json:"chain_id"`
	TokenId      string   `json:"token_id"`
	Owner        string   `json:"owner"`
	URI          string   `json:"uri"`
	Name         string   `json:"name"`
	Symbol       string   `json:"symbol"`
	Contract     string   `json:"contract"`
	ContractType Contract `json:"contract_type"`
}

var ErrIncorrectContractType = errors.New("incorrect contract type for nft. must be ERC721 or ERC1155")

func NewEthNft(chainId, tokenId, owner, uri, name, symbol, contract string, contractType Contract) (EthNft, error) {
	if !(contractType == ERC721 || contractType == ERC1155) {
		return EthNft{}, ErrIncorrectContractType
	}
	return EthNft{
		ChainId:      chainId,
		TokenId:      tokenId,
		Owner:        owner,
		URI:          uri,
		Name:         name,
		Symbol:       symbol,
		Contract:     contract,
		ContractType: contractType,
	}, nil
}
