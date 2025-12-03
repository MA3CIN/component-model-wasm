package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[Go] Received request...")

		// Outgoing requests work natively in TinyGo wasip2
		resp, err := http.Get("http://127.0.0.1:8081/getData")
		if err != nil {
			http.Error(w, "Failed to call C service: "+err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		fmt.Fprintf(w, "Go Service received: %s", string(body))
	})

	// Instead of http.ListenAndServe, use the adapter's Serve function.
	// This satisfies the 'wasi:http/incoming-handler' world requirements.
	// wasihttp.Serve(http.DefaultServeMux)

	// TODO: this is broken lol
}
