package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/readings", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("http://localhost:8082/readings")
		if err != nil {
			http.Error(w, "DB Service Down", 500)
			return
		}
		defer res.Body.Close()

		io.Copy(w, res.Body)
	})

	log.Println("API Gateway Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
