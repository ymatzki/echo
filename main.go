package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("start http server...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("access: %+v\n", r)
	resp := response{Message: os.Getenv("MESSAGE")}
	j, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
