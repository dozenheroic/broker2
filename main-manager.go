package main

import (
	"encoding/json"
	"message-broker/manager"
	"net/http"
)

var state = manager.LoadState()

func status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(state)
}

func main() {
	http.HandleFunc("/status", status)

	println("Manager running on :8090")
	http.ListenAndServe(":8090", nil)
}
