package main

import (
	"log"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	http.HandleFunc("/ping", pingHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
