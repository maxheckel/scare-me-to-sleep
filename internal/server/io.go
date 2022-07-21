package server

import (
	"encoding/json"
	"net/http"
)

type ErrOut struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (a App) writeError(err error, w http.ResponseWriter) {
	errOut := ErrOut{
		Error: err.Error(),
		Code:  500,
	}
	w.WriteHeader(errOut.Code)
	outBytes, err := json.Marshal(errOut)
	if err != nil {
		panic(err)
	}
	w.Write(outBytes)
}

func (a App) writeJSON(obj interface{}, w http.ResponseWriter) {
	outBytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Write(outBytes)
}
