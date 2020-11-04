//go:generate ${GOPATH}/bin/rice embed-go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	var in io.Reader = os.Stdin
	var err error
	var lines = make([]string, 0, 10)
	reader := bufio.NewReaderSize(in, 4096)
	for line := ""; err == nil; line, err = reader.ReadString('\n') {
		lines = append(lines, line)
	}
	fmt.Print(lines[0])
	mapB, _ := json.Marshal(lines)
	w.Write([]byte(mapB))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/logs", LogsHandler)
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
	// http.Handle("/", http.FileServer(rice.MustFindBox("http-files").HTTPBox()))
	address := "0.0.0.0:8000"
	fmt.Printf("Starting http server on %s\n", address)
	http.ListenAndServe(address, r)
}
