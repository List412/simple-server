package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"simple-server/internal/config"
	"simple-server/internal/handlers"
	"time"
)

func Start(ctx context.Context, cfg config.HttpServer) error {
	router := mux.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	err := handlers.ApplyHandlers(router)
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)

	fmt.Println(fmt.Sprintf("listening on  :%s", cfg.Port))
	go serve(srv, errChan)

	select {
	case <-ctx.Done():
		fmt.Println("shutdown server")
		_ = srv.Shutdown(ctx)
		return nil
	case err := <-errChan:
		fmt.Println(err.Error())
	}
	return nil
}

func serve(srv *http.Server, errChan chan error) {
	err := srv.ListenAndServe()
	errChan <- err
}
