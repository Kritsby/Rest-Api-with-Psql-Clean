package router

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	http2 "rest-api/internal/controller/http"
	"time"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	return router
}

func NewConnection(handler *http2.FirstHandler) {

	var wait time.Duration = time.Second * 30

	router := httprouter.New()

	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	handler.Register(router)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutdown server")
	os.Exit(0)
}
