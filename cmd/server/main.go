package main

import (
	"fmt"
	"net/http"

	"github.com/praneshkumarc/snapvault.io/internal/api"
)

func main() {
	mux := api.NewRouter()

	fmt.Println("SnapVault Server listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
