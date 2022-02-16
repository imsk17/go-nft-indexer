package domain

type EthNftID struct {
	Contract string `json:"contract"`
	TokenId  string `json:"tokenId"`
	Owner    string `json:"owner"`
	Type     string `json:"type"`
}
