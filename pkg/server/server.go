package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	av "github.com/preflightsiren/stocks/pkg/alphavantage"
)

type Server struct {
	Port   int
	Ndays  int
	Symbol string
    Client *av.Client
}

func NewServer(Port, Ndays int, Symbol string) *Server {
    client := av.NewClient()
	return &Server{
		Port:   Port,
		Ndays:  Ndays,
		Symbol: Symbol,
        Client: client,
	}
}

func (s *Server) Start() {
    mux := http.NewServeMux()

	h1 := func(w http.ResponseWriter, _ *http.Request) {
        resp, err := s.Client.GetSymbol(s.Symbol)
        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            io.WriteString(w, err.Error())
        }
        resp_string, err := json.Marshal(resp.Data)
		io.WriteString(w, string(resp_string))
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "OK\n")
	}

	mux.HandleFunc("/", h1)
	mux.HandleFunc("/healthz", h2)
    log.Printf("Starting server with ndays: %d symbol: %s", s.Ndays, s.Symbol)
    log.Printf("Listening on %d", s.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",s.Port), mux))

}
