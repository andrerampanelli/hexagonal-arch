package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrerampanelli/hexagonal-arch/adapters/web/handler"
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service interfaces.ProductServiceInterface
}

func NewWebServer(service interfaces.ProductServiceInterface) *WebServer {
	return &WebServer{Service: service}
}

func (s *WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, s.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
