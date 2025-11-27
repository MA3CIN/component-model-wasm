package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Go] Received request on /fetch, calling C service...")

		resp, err := http.Get("http://127.0.0.1:8081/getData")
		if err != nil {
			http.Error(w, "Failed to call C service: "+err.Error(), 500)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		
		fmt.Fprintf(w, "Go Service received from C Service: %s", string(body))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("[Go] Server starting on http://127.0.0.1:%s\n", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}