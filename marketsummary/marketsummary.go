package marketsummary

// Example of the resulting market data output
// {"market":5775,"total_volume":1234567.89,"mean_price":23.33,"mean_volume":6144.299,"volume_weighted_average_price":5234.2,"percentage_buy":0.50}

// Example of trade data
// {"id":50714,"market":2782,"price":26.602724094526256,"volume":3077.653731797662,"is_buy":true}

type MarketSummary struct {
	Market          int
	TotalVolume     float64
	CumulativePrice float64
	AmountOfPrices  int
	MeanPrice       float64
	MeanVolume      float64
}

type Trade struct {
	Id     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

// CalculateMarketMeanPrice calculates the mean price per market
func CalculateMarketMeanVolume(marketsSummary map[int]*MarketSummary, t Trade) {
	if _, ok := marketsSummary[t.Market]; ok {
		marketSummary := marketsSummary[t.Market]
		marketSummary.MeanVolume = marketSummary.TotalVolume / float64(marketSummary.AmountOfPrices)
	}
}

// CalculateMarketMeanPrice calculates the mean price per market
func CalculateMarketMeanPrice(marketsSummary map[int]*MarketSummary, t Trade) {
	if _, ok := marketsSummary[t.Market]; ok {
		marketSummary := marketsSummary[t.Market]
		marketSummary.CumulativePrice += t.Price
		marketSummary.AmountOfPrices += 1
		marketSummary.MeanPrice = marketSummary.CumulativePrice / float64(marketSummary.AmountOfPrices)
	}
}

// CalculateMarketVolume calculates the running total of the volume per market
func CalculateMarketVolume(marketsSummary map[int]*MarketSummary, t Trade) {
	if _, ok := marketsSummary[t.Market]; ok {
		marketsSummary[t.Market].TotalVolume = marketsSummary[t.Market].TotalVolume + t.Volume
	} else {
		marketsSummary[t.Market] = &MarketSummary{Market: t.Market, TotalVolume: t.Volume}
	}
}
