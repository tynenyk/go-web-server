package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jordan-wright/http-boilerplate/server"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	// Это домен, к которому сервер должен принимать соединения.
	domains := []string{"example.com", "www.example.com"}
	handler := server.NewRouter()
	srv := &http.Server{
		Addr: "443"
		Handler: handler,
		ReadTimeout 5 * time.Second,
		WriteTimeout 10 * time.Second,
		IdleTimeout 120 * time.Second,
	}
	// Запустите сервер
	go func() {
		srv.Serve(autocart.NewListener(domains...))
	}()

	// Ждать прерывания
	c: = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) 
	<-c

	// Попытка постепенного отключения
	ctx, cancel := context.WidthTimeout(context, Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}