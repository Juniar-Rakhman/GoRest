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
	SetDefaultHeader(w, 200)
	output := SetFormat(repo.FindAll())
	fmt.Fprintln(w, string(output))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productId"])

	SetDefaultHeader(w, 200)
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
		SetDefaultHeader(w, 422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	SetDefaultHeader(w, http.StatusCreated)
	output := SetFormat(repo.CreateProduct(prod))
	fmt.Fprintln(w, string(output))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productId"])

	SetDefaultHeader(w, 200)
	output := repo.DeleteProduct(id)
	fmt.Fprintln(w, string(output))
}

//---- Cart Handlers ----//

func AddItemToNewCart(w http.ResponseWriter, r *http.Request) {

}

func AddItemToExistingCart(w http.ResponseWriter, r *http.Request) {

}

func GetExistingCart(w http.ResponseWriter, r *http.Request) {

}

func DeleteItemFromCart(w http.ResponseWriter, r *http.Request) {

}

//---- Payment Handlers ----//
func GetPaymentDetails(w http.ResponseWriter, r *http.Request) {

}

func SetCartToPaid(w http.ResponseWriter, r *http.Request) {

}

func AddDiscountCode(w http.ResponseWriter, r *http.Request) {

}