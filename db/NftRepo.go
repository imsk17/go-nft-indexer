package db

import (
	"context"
	"go-nft-listener/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NftRepoReadService interface {
	GetNftInfo(contract, tokenId string) (domain.EthNftID, error)
	GetNftsByOwner(owner string) ([]domain.EthNftID, error)
}

type NftRepoWriteService interface {
	CreateOrUpdateFull(info domain.EthNftID) error
	CreateOrUpdateOwner(info domain.EthNftID) error
}

type nftSvc struct {
	db *mongo.Collection
}

func NewNftRepo(db *mongo.Collection) NftRepoReadService {
	return &nftSvc{db: db}
}

func (s *nftSvc) GetNftInfo(contract, tokenId string) (domain.EthNftID, error) {
	var info domain.EthNftID
	if err := s.db.FindOne(context.Background(), bson.M{"contract": contract, "tokenId": tokenId}).Decode(&info); err != nil {
		return domain.EthNftID{}, err
	}
	return info, nil
}

func (s *nftSvc) GetNftsByOwner(owner string) ([]domain.EthNftID, error) {
	var info []domain.EthNftID
	cursor, err := s.db.Find(context.Background(), bson.M{"owner": owner})
	if err != nil {
		return info, err
	}
	for cursor.Next(context.Background()) {
		var i domain.EthNftID
		if err := cursor.Decode(&i); err != nil {
			return info, err
		}
		info = append(info, i)
	}
	return info, nil
}

func (s *nftSvc) CreateOrUpdateOwner(info domain.EthNft) error {
	if err := s.db.FindOne(context.Background(), bson.M{"contract": info.Contract, "tokenId": info.TokenId}).Err(); err == mongo.ErrNoDocuments {
		_, err := s.db.InsertOne(context.Background(), info)
		return err
	}
	_, err := s.db.UpdateOne(context.Background(), bson.M{"contract": info.Contract, "tokenId": info.TokenId}, info)
	return err
}

func (s *nftSvc) CreateOrUpdateFull(info domain.EthNft) error {
	if err := s.db.FindOne(context.Background(), bson.M{"contract": info.Contract, "tokenId": info.TokenId}).Err(); err == mongo.ErrNoDocuments {
		_, err := s.db.InsertOne(context.Background(), info)
		return err
	}
	_, err := s.db.UpdateOne(context.Background(), bson.M{"contract": info.Contract, "tokenId": info.TokenId}, info)
	return err
}
