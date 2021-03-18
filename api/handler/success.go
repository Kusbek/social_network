package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func successResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
