test:
	go test -v ./pkg/...

clean:
	rm stocks

build: clean
	go build -o stocks cmd/stocks.go

update-data:
	curl "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY&symbol=MSFT" > ./pkg/alphavantage/stock.json

.PHONY: test build clean update-data
