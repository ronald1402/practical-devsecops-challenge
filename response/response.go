package response

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, data interface{}, statusCode int) {
	bytesResponse, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytesResponse)
}

func ResponseByte(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
