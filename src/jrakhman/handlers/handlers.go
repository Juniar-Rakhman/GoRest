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
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetProductAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(productList); err != nil {
		panic(err)
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productId"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(RepoFindProduct(id)); err != nil {
		panic(err)
	}
}

func GetProductByCategory(w http.ResponseWriter, r *http.Request) {

}

func GetProductByPrice(w http.ResponseWriter, r *http.Request) {

}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var prod model.Product

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

	if err := json.NewEncoder(w).Encode(RepoCreateProduct(prod)); err != nil {
		panic(err)
	}
}

func DelProduct(w http.ResponseWriter, r *http.Request) {

}