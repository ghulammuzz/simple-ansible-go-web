package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hallo gaes",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Lima Ratus", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/hehe", helloHandler)
	http.ListenAndServe(":5000", nil)
}
