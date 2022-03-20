package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	ms "messari-challenge/marketsummary"
	"os"
	"strings"
)

// type MarketSummary map[int]TotalMarketVolume

func main() {
	marketsSummary := make(map[int]*ms.MarketSummary)
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

		var trade ms.Trade
		if err := json.Unmarshal([]byte(line), &trade); err != nil {
			log.Fatalf("some unknown error occurred %v", err)
		}

		ms.CalculateMarketVolume(marketsSummary, trade)
		ms.CalculateMarketMeanPrice(marketsSummary, trade)

		// if trade.Id == 100 {
		// 	break
		// }
	}

	encoder := json.NewEncoder(os.Stdout)
	for _, s := range marketsSummary {
		encoder.Encode(s)
	}
}
