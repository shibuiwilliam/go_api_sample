package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	port := os.Getenv("PORT")
	httpPort := fmt.Sprintf(":%v", port)
	log.Printf("port: %v", httpPort)

	router := mux.NewRouter().PathPrefix("/v1").Subrouter()
	router.HandleFunc("/ping", pingHandler).Methods("GET", "POST")
	server := &http.Server{
		Handler: router,
		Addr:    httpPort,
	}
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
