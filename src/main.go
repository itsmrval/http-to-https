package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetURL := "https://" + r.Host + r.RequestURI
		log.Printf("[REDIRECT] %s - %s (%s)", targetURL, r.RemoteAddr, r.UserAgent())
		http.Redirect(w, r, targetURL, http.StatusMovedPermanently)
	})

	log.Printf("[INFO] Service running on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}
