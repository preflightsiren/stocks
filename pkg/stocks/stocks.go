package stocks

import av "github.com/preflightsiren/stocks/pkg/alphavantage"

type Result struct {
    Items []av.Day
    Average int
    NDays int
}


