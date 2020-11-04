//go:generate ${GOPATH}/bin/rice embed-go
package main

import (
	"fmt"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
	// http.Handle("/", http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
	address := "0.0.0.0:8000"
	fmt.Printf("Starting http server on %s\n", address)
	http.ListenAndServe(address, r)
}
