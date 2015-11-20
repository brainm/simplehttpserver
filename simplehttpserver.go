package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

var (
	root string
)

func init() {
	root, _ = filepath.Abs(filepath.Dir(os.Args[0]))
}

func main() {
	host := flag.String("host", "localhost", "http host")
	port := flag.Int("port", 8000, "http port")

	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(root)))

	listen := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Listening %s...", listen)
	http.ListenAndServe(listen, router)
}
