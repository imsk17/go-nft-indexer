package domain

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Contract string

const (
	ERC721  Contract = "ERC721"
	ERC1155 Contract = "ERC1155"
)

type EthNft struct {
	ChainId      string    `json:"chain_id"`
	TokenId      string    `json:"token_id"`
	Owner        string    `json:"owner"`
	URI          string    `json:"uri"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Contract     Contract  `json:"contract"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ContractType string    `json:"contract_type"`
}

func NewEthNft(chainId, tokenId, owner, uri, name, symbol, contractType string, contract Contract) EthNft {
	if !(contract == "ERC721" || contract == "ERC1155") {
		log.Panic("contract must be ERC721 or ERC1155")
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
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
