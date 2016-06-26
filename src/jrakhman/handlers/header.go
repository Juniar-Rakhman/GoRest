package handlers

import (
	"encoding/xml"
	"encoding/json"
	"net/http"
)

var Format string

func SetFormat(data interface{}) []byte {

	var apiOutput []byte

	Format = "json" //change accordingly

	if Format == "json" {
		output, _ := json.Marshal(data)
		apiOutput = output
	} else if Format == "xml" {
		output, _ := xml.Marshal(data)
		apiOutput = output
	}

	return apiOutput
}

func SetDefaultHeader(w http.ResponseWriter, status int) {

	w.Header().Set("Allow", "GET,OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
}