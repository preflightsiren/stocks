FROM docker.io/library/golang:1.20-bookworm AS build


WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.* ./
RUN go mod download && go mod verify

COPY . .

RUN go build --ldflags '-w -s -linkmode external -extldflags "-static"' \
-asmflags -trimpath -v -o stocks cmd/stocks.go

FROM scratch
COPY --from=build /usr/src/app/stocks ./

ENV NDAYS 7
ENV SYMBOL MSFT
EXPOSE 8080

CMD ["./stocks"]
