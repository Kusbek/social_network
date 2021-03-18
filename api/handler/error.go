package handler

import (
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, err error) {
	log.Println(err.Error())
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}
