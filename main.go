package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-NetworkManager-Status", "online")
		w.WriteHeader(http.StatusNoContent)
	})
	log.Println("Server is running!")
	http.ListenAndServe(":80", nil)
}
