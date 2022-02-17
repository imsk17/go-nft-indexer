package db

import (
	"errors"
	"go-nft-listener/domain"
	"go-nft-listener/events"

	"gorm.io/gorm"
)

type NftRepoReadService interface {
	GetNftInfo(contract, tokenId string) (domain.EthNft, error)
	GetNftsByOwner(owner string, chainId string) ([]domain.EthNft, error)
}

type NftRepoWriteService interface {
	CreateOrUpdateFull(info events.Event) error
	CreateOrUpdateOwner(info events.Event) error
}

type NftRepoService interface {
	NftRepoReadService
	NftRepoWriteService
}

func NewNftRepo(db *gorm.DB) NftRepoService {
	return &nftSvc{db: db}
}

type nftSvc struct {
	db *gorm.DB
}

func (n *nftSvc) GetNftInfo(contract, tokenId string) (domain.EthNft, error) {
	var nft domain.EthNft
	if err := n.db.First(&nft, "contract = ? AND tokenId = ?", contract, tokenId).Error; err != nil {
		return nft, err
	}
	return nft, nil
}

func (n *nftSvc) GetNftsByOwner(owner string, chainId string) ([]domain.EthNft, error) {
	var nfts []domain.EthNft
	if err := n.db.Find(&nfts, "owner = ? AND chainId = ?", owner, chainId).Error; err != nil {
		return nfts, err
	}
	return nfts, nil
}

func (n *nftSvc) CreateOrUpdateFull(info events.Event) error {
	return errors.New("not implemented")
}

func (n *nftSvc) CreateOrUpdateOwner(info events.Event) error {
	return errors.New("not implemented")
}
