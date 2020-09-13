package wrapper

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Error   string      `json:"error"`
	Message interface{} `json:"msg"`
}

func ErrorResponse(w http.ResponseWriter, err error) {
	data, err := json.Marshal(response{
		Error:   err.Error(),
		Message: nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	http.Error(w, string(data), http.StatusBadRequest)
}

func Response(w http.ResponseWriter, msg interface{}) {
	if err := json.NewEncoder(w).Encode(response{Error: "", Message: msg}); err != nil {
		ErrorResponse(w, err)
	}
}
