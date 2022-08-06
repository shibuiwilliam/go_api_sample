package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sys/unix"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// ping := Ping{http.StatusOK, "ok"}

	// res, err := json.Marshal(ping)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// log.Println("OK; GREAT!")

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(res)
}

func pingErrorHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AAAAAAAAAAAAAAAAAAAAAA")
	http.Error(w, errors.New("AAAAAAAAAAAAA").Error(), http.StatusInternalServerError)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	port := os.Getenv("PORT")
	httpPort := fmt.Sprintf(":%v", port)
	log.Printf("port: %v", httpPort)

	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET", "POST")
	router.HandleFunc("/pingE", pingErrorHandler).Methods("GET", "POST")
	server := &http.Server{
		Handler: router,
		Addr:    httpPort,
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, unix.SIGTERM)
	select {
	case <-sigCh:
		log.Println("received SIGTERM, exiting gracefully")
	case <-ctx.Done():
	}

	log.Println("gracefully shutdown http server")
	if err := server.Shutdown(ctx); err == nil {
		log.Println("completed http server graceful shutdown")
	} else {
		log.Fatalf("failed to shutdown http server", err)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		log.Fatalf("unhandled error received", err)
	}

	log.Println("exiting with 0")
}
