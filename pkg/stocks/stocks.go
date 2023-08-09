package stocks

import (
	"strconv"
    "math"
	av "github.com/preflightsiren/stocks/pkg/alphavantage"
)

type Result struct {
	Items   *[]av.Day
	Average float64
	NDays   int
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

func AverageResults(input *av.Result, ndays int) *Result {
	var index []string
	for k := range input.Data {
		index = append(index, k)
	}
	var items []av.Day
	for i := 0; i < ndays; i++ {
		items = append(items, input.Data[index[i]])
	}
    var sum float64
    for v := range(items) {
        f, err := strconv.ParseFloat(items[v].Close, 64)
        if err != nil {
            continue
        }
        sum = sum + f
    }
    average := roundFloat(sum / float64(ndays), 4)

	return &Result{
		Items:   &items,
		Average: average,
		NDays:   ndays,
	}
}
