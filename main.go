package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/aimtest", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
	})
	http.ListenAndServe(":8080", nil)
}
