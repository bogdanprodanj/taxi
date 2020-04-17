package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/valyala/fasthttp"
)

var quit = make(chan struct{})

func main() {
	s := &server{storage: newStorage()}
	go s.storage.cancel()
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/request":
			s.request(ctx)
		case "/admin/requests":
			s.adminRequest(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	s.srv = &fasthttp.Server{Handler: m}
	log.Println("listening on port 8080")
	go s.listenAndServe()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	log.Println("shutting down")
	s.srv.Shutdown()
	close(quit)
}

type server struct {
	srv     *fasthttp.Server
	storage *storage
}

func (s *server) listenAndServe() {
	err := s.srv.ListenAndServe("127.0.0.1:8080")
	if err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func (s *server) request(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, s.storage.makeRequest())
}

func (s *server) adminRequest(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, s.storage.getAllRequests())
}
