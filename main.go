package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

// {"id":50714,"market":2782,"price":26.602724094526256,"volume":3077.653731797662,"is_buy":true}

type Trade struct {
	Id     int     `json:"id"`
	Market int     `json:"market"`
	Price  float64 `json:"price"`
	Volume float64 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

// {"market":5775,"total_volume":1234567.89,"mean_price":23.33,"mean_volume":6144.299,"volume_weighted_average_price":5234.2,"percentage_buy":0.50}
type MarketSummary struct {
	Market      int
	TotalVolume float64
}

// type MarketSummary map[int]TotalMarketVolume

func main() {
	marketsSummary := make(map[int]*MarketSummary)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			log.Println("no data")
			break
		} else if err != nil {
			log.Fatalf("some unknown error occurred %v", err)
		}

		if strings.TrimSpace(line) == "BEGIN" {
			continue
		}

		if strings.TrimSpace(line) == "END" {
			log.Println("End of file reached!")
			break
		}

		var trade Trade
		if err := json.Unmarshal([]byte(line), &trade); err != nil {
			log.Fatalf("some unknown error occurred %v", err)
		}

		CalculateMarketVolume(marketsSummary, trade)

		if trade.Id == 100 {
			break
		}
	}

	encoder := json.NewEncoder(os.Stdout)
	for _, s := range marketsSummary {
		encoder.Encode(s)
	}
}

// CalculateMarketVolume calculates the running total of the volume per market
func CalculateMarketVolume(marketSummary map[int]*MarketSummary, t Trade) {
	if _, ok := marketSummary[t.Market]; ok {
		marketSummary[t.Market].TotalVolume = marketSummary[t.Market].TotalVolume + t.Volume
	} else {
		marketSummary[t.Market] = &MarketSummary{Market: t.Market, TotalVolume: t.Volume}
	}
}
