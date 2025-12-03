package main

import (
	"net/http"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net/wasihttp"
)

func init() {
	wasihttp.HandleFunc(handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	logger := wasilog.ContextLogger("handler")

	logger.Info("request received", "host", r.Host, "path", r.URL.Path, "agent", r.Header.Get("User-Agent"))

	_, err := w.Write([]byte("hello world!"))
	if err != nil {
		logger.Error("failed to write body", "error", err)
	}
}

func main() {}
