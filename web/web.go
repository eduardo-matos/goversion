package web

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/eduardo-matos/goversion/vchecker"
	"github.com/gorilla/mux"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Port number")
	flag.Parse()
}

// Run runs web version checker
func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/{version}", versionHandler)

	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version := vars["version"]

	valid, err := vchecker.Version(version, nil)
	vchecker.Output(w, valid, err)
}
