package main

import (
	"context"
	"hello_world/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//log to stdout
	l := log.New(os.Stdout, "hello-world", log.LstdFlags)

	//create handlers for various endpoints
	h := handlers.NewHello(l)
	ph := handlers.NewProducts(l)
	p := handlers.NewProduct(l)

	//create new servermux to server our endpoints
	sm := http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/product/", p)
	sm.Handle("/products", ph)

	//create server with appropriate settings
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//run server as a separate go routine so that it doesn't block main
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//create channel to keep track of server kill/interrupt attempt
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	<-sigChan
	l.Println("Attempting shutdown")

	//give 30 seconds to handle existing requests and then shutdown (gracefully)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
