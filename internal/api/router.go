package api

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("POST /v1/backups/init", handleBackupInit)
	mux.HandleFunc("POST /v1/backups/complete", handleBackupComplete)
	return mux
}
