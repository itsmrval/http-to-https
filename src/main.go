package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetURL := "https://" + r.Host + r.RequestURI
		http.Redirect(w, r, targetURL, http.StatusMovedPermanently)
	})

	log.Println("Service running on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
