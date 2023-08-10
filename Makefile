MAKEFLAGS += --silent
PODMAN := $(shell which podman)

test:
	go test -v ./pkg/...

clean:
	rm stocks

build:
	go build -o stocks cmd/stocks.go

run:
	go run cmd/stocks.go

build-container:
	$(PODMAN) build -t stocks:latest -f Dockerfile

NDAYS = 7
SYMBOL = MSFT
run-container:
	$(PODMAN) run -e NDAYS=$(NDAYS) -e SYMBOL=$(SYMBOL) -p 8080:8080 stocks:latest

update-data:
	curl "https://www.alphavantage.co/query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY&symbol=MSFT" > ./pkg/alphavantage/stock.json

.PHONY: test build clean run build-container run-container update-data prereqs
