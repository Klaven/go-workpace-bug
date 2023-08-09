package main

import (
	"fmt"
	"net/http"
	"os"

	log2 "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	pb "github.com/null-channel/go-workpace-bug/proto"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	logger := log2.New(os.Stderr)
	if true {
		logger.Warn("chewy!", "butter", true)
	}
	fmt.Println("not serving it. just return from app a")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ProdutsHandler")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	c := pb.CreateApplicationReply{}
	fmt.Println(c)
}
