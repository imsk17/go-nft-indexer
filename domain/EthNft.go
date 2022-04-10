package domain

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

type Contract string

const (
	ERC721  Contract = "ERC721"
	ERC1155 Contract = "ERC1155"
)

type EthNft struct {
	ID           uint     `gorm:"primarykey"`
	ChainId      string   `json:"chain_id" gorm:"index:idx_chain_id_owner"`
	TokenId      string   `json:"token_id"`
	Owner        string   `json:"owner" gorm:"index:idx_chain_id_owner"`
	URI          string   `json:"uri"`
	Name         string   `json:"name"`
	Symbol       string   `json:"symbol"`
	Contract     string   `json:"contract"`
	ContractType Contract `json:"contract_type"`
	UpdatedAt    int64    `gorm:"autoUpdateTime:milli"`
}

var ErrIncorrectContractType = errors.New("incorrect contract type for nft. must be ERC721 or ERC1155")

func NewEthNft(chainId, tokenId, owner, uri, name, symbol, contract string, contractType Contract) (EthNft, error) {
	if !(contractType == ERC721 || contractType == ERC1155) {
		return EthNft{}, ErrIncorrectContractType
	}
	ownerChecksum := common.HexToAddress(owner).String()
	contractChecksum := common.HexToAddress(contract).String()
	return EthNft{
		ChainId:      chainId,
		TokenId:      tokenId,
		Owner:        ownerChecksum,
		URI:          uri,
		Name:         name,
		Symbol:       symbol,
		Contract:     contractChecksum,
		ContractType: contractType,
	}, nil
}
