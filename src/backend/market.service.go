package backend

import (
	"crypto-market-terminal/src/backend/client"
	"crypto-market-terminal/src/backend/models"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type MarketService struct {
	clientHttp *client.ClientHTTP
	config     *Config
}

type Response struct {
	Data []models.Currency `json:"data"`
}

func NewMarketService(clientHttp *client.ClientHTTP, config *Config) *MarketService {
	return &MarketService{
		clientHttp,
		config,
	}
}

func (m *MarketService) RetrieveCryptoRate() *[]models.Currency {
	var response Response
	resp := m.clientHttp.Get(fmt.Sprintf("https://api.coincap.io/v2/assets?limit=%v", m.config.amountCurrency), m.config.token)
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&response)
	if err != nil {
		log.Panicln("Error in decode data response")
	}
	mappingData(&response.Data)
	return &response.Data
}

func mappingData(data *[]models.Currency) {
	for i := range *data {
		price, _ := strconv.ParseFloat((*data)[i].Price, 64)
		volume, _ := strconv.ParseFloat((*data)[i].Volume24H, 64)
		percent24h, _ := strconv.ParseFloat((*data)[i].ChangePercent24H, 64)
		(*data)[i].ChangePercent24H = roundedDecimals(percent24h)
		(*data)[i].Price = roundedDecimals(price)
		(*data)[i].Volume24H = roundedDecimals(volume)
	}
}

func roundedDecimals(data float64) string {
	return fmt.Sprintf("%.2f", data)
}
