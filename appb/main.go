package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	pb "github.com/null-channel/go-workpace-bug/proto"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	fmt.Println("not serving it. just return from app b")
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
