package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("URL")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("OH NO %v!", err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("OH NO %v!", err)
	}
	log.Printf("AAAAAAAAAAAA %v \n", string(byteArray))

	ping := Ping{http.StatusOK, "ok"}

	res, err := json.Marshal(ping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/ping", pingHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
