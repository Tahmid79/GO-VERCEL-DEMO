package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	msg := response{"Hello from go"}
	json.NewEncoder(w).Encode(&msg)
	//fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
