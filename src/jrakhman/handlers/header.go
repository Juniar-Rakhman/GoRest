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

func SetDefaultHeader(w http.ResponseWriter) http.ResponseWriter {

	w.Header().Set("Allow", "GET,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	return w
}