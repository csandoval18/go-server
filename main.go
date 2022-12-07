package main

import (
	"net/http"

	"github.com/csandoval18/go-server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":4000", srv)
}
