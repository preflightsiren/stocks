package main

import (
	"log"
	"os"
	"strconv"

	"github.com/preflightsiren/stocks/pkg/server"
)

func main() {
    var ndays int // Number of days to return data for
    var symbol string // Financial symbol to query from alphavantage
    if os.Getenv("NDAYS") == "" {
        ndays = 7
    } else {
        var err error
        ndays, err  = strconv.Atoi(os.Getenv("NDAYS"))
        if err != nil {
            log.Fatal(err)
        }
    }
    if os.Getenv("SYMBOL") == "" {
        symbol = "MSFT"
    } else {
        symbol = os.Getenv("SYMBOL")
    }
    server := server.NewServer(8080, ndays, symbol)
    server.Start()

}
