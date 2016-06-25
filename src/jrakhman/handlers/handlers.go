package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"io/ioutil"
	"io"
	"jrakhman/model"
	"jrakhman/repo"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to GoRestAPI!")
}

func GetProductAll(w http.ResponseWriter, r *http.Request) {
	SetDefaultHeader(w)
	output := SetFormat(repo.FindAll())
	fmt.Fprintln(w, string(output))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productId"])

	SetDefaultHeader(w)
	output := SetFormat(repo.FindProduct(id))
	fmt.Fprintln(w, string(output))
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var prod model.Product

	//limit given json
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &prod); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(repo.CreateProduct(prod)); err != nil {
		panic(err)
	}
}

func DelProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productId"])

	SetDefaultHeader(w)
	output := SetFormat(repo.DeleteProduct(id))
	fmt.Fprintln(w, string(output))
}