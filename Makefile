test:
	go test -v ./pkg/...

update-data:
	curl "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY&symbol=MSFT" > ./pkg/alphavantage/stock.json

.PHONY: test update-data
