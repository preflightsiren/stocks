package stocks

import (
	"math"
	"sort"
	"strconv"

	av "github.com/preflightsiren/stocks/pkg/alphavantage"
)

type Result struct {
	Items   *[]av.Day
	Average float64
	NDays   int
}

// rounds floats to the nearest precision
// eg. 1.23456790 , 3 returns 1.235
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func AverageResults(input *av.Result, ndays int) *Result {
	var index []string
	// build a lookup-index of dates we have data for
	for k := range input.Data {
		index = append(index, k)
	}
	// sort the index so we guarantee the order of the dates
	sort.Strings(index)
	var items []av.Day
	// Starting with the end of the array (latest date),
	// count back ndays from the index getting the last ndays dates
	for i := len(index) - 1; i > len(index)-ndays-1; i-- {
		items = append(items, input.Data[index[i]])
	}
	// Calculate the sum of the daily close values
	var sum float64
	for v := range items {
		f, err := strconv.ParseFloat(items[v].Close, 64)
		// if the data is malformed, do not include in the sum.
		if err != nil {
			continue
		}
		sum = sum + f
	}
	// Round float to prevent extra precision
	average := roundFloat(sum/float64(ndays), 4)

	return &Result{
		Items:   &items,
		Average: average,
		NDays:   ndays,
	}
}
