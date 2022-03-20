package marketsummary_test

import (
	"fmt"
	"messari-challenge/marketsummary"
	"testing"
)

func TestCalculateMarketVolume(t testing.T) {
	t.Parallel()

	trades := []marketsummary.Trade{
		{Id: 1, Market: 123, Price: 10, Volume: 123, IsBuy: false},
		{Id: 1, Market: 123, Price: 10, Volume: 555.26, IsBuy: false},
		{Id: 1, Market: 123, Price: 10, Volume: 600.5, IsBuy: false},
		{Id: 1, Market: 123, Price: 10, Volume: 9, IsBuy: false},
	}

	fmt.Printf("trades: %v\n", trades)
}

func TestCalculateMarketMeanPrice(t testing.T) {
	t.Parallel()
	marketsSummary := make(map[int]*marketsummary.MarketSummary)

	trades := []marketsummary.Trade{
		{Id: 1, Market: 1, Price: 1.2, Volume: 0, IsBuy: false},
		{Id: 2, Market: 1, Price: 1.2, Volume: 0, IsBuy: false},
		{Id: 3, Market: 1, Price: 1.2, Volume: 0, IsBuy: false},
		{Id: 4, Market: 4, Price: 1.2, Volume: 0, IsBuy: false},
	}

	fmt.Printf("trades: %v\n", trades)
	fmt.Printf("marketsSummary: %v\n", marketsSummary)
}
