package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/piquette/finance-go/equity"
)

func getSymbolInfo(symbols string) string {
	var result string
	log.Println("symbols given: ", symbols)
	symbolList := strings.Split(symbols, ",")
	if len(symbolList) == 0 || symbolList[0] == "" {
		result = "Please name SYMBOLS to your preferred stock ticker names, comma seperated. Ie AAPL,TSLA"
		return result
	}
	iter := equity.List(symbolList)
	// Iterate over results. Will exit upon any error.
	for iter.Next() {
		q := iter.Equity()
		result = result + fmt.Sprintf("%s (%s): Bid: %.2f Ask: %.2f Price: %.2f High: %.2f Low: %.2f Close: %.2f Post: %.2f Currency: %s Market State: %s\n",
			q.Symbol,
			q.ShortName,
			q.Bid,
			q.Ask,
			q.RegularMarketPrice,
			q.RegularMarketDayHigh,
			q.RegularMarketDayLow,
			q.RegularMarketPreviousClose,
			q.RegularMarketPrice+q.PreMarketChange,
			q.CurrencyID,
			q.MarketState)
	}
	log.Println(result)
	return result
}