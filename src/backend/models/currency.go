package models

type Currency struct {
	Rank             string `json:"rank"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Price            string `json:"priceUsd"`
	ChangePercent24H string `json:"changePercent24Hr"`
	Volume24H        string `json:"volumeUsd24Hr"`
}
