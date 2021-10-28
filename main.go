package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greeting/{name}", greetingHandler)
	http.ListenAndServe(":8080", r)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resp := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("hello, %s", vars["name"]),
	}
	json.NewEncoder(w).Encode(&resp)
}
