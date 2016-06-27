package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"jrakhman/router"
	"github.com/gorilla/mux"
	"encoding/json"
	"bytes"
)

//Functional testing is covered here. For code coverage testing refer to handlers_test.go

var (
	muxRouter *mux.Router
	respRec *httptest.ResponseRecorder
)

func setup() {
	muxRouter = router.NewRouter()

	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

//func TestGetIndex(t *testing.T) {
//	setup()
//
//	method := "GET"
//	urlStr := "/"
//
//	req, err := http.NewRequest(method, urlStr, nil)
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//
//	muxRouter.ServeHTTP(respRec, req)
//
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
//	}
//}
//
//func TestGetProductAll(t *testing.T) {
//	setup()
//
//	method := "GET"
//	urlStr := "/product"
//
//	req, err := http.NewRequest(method, urlStr, nil)
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//
//	muxRouter.ServeHTTP(respRec, req)
//
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
//	}
//}
//
//func TestGetProductId(t *testing.T) {
//	setup()
//
//	method := "GET"
//	urlStr := "/product/1"
//
//	req, err := http.NewRequest(method, urlStr, nil)
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//
//	muxRouter.ServeHTTP(respRec, req)
//
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
//	}
//}
//
//func TestAddProduct(t *testing.T) {
//	setup()
//
//	method := "POST"
//	urlStr := "/product"
//
//	postBody := map[string]interface{}{
//		"Name": "Testing3",
//		"Size": 99,
//		"Color":"magenta",
//		"Price":100,
//	}
//	body, _ := json.Marshal(postBody)
//	req, err := http.NewRequest(method, urlStr, bytes.NewReader(body))
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//	muxRouter.ServeHTTP(respRec, req)
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusCreated {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusCreated)
//	}
//}
//
//func TestDeleteProduct(t *testing.T) {
//	setup()
//
//	method := "DELETE"
//	urlStr := "/product/10"
//
//	req, err := http.NewRequest(method, urlStr, nil)
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//
//	muxRouter.ServeHTTP(respRec, req)
//
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusOK)
//	}
//}
//
//func TestAddItemToNewCart(t *testing.T) {
//	setup()
//
//	method := "POST"
//	urlStr := "/cart/new/423"
//
//	postBody := map[string]interface{}{
//		"prodId": 3,
//		"qty": 33,
//		"price":100,
//	}
//
//	body, _ := json.Marshal(postBody)
//	req, err := http.NewRequest(method, urlStr, bytes.NewReader(body))
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//	muxRouter.ServeHTTP(respRec, req)
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusCreated {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusCreated)
//	}
//}
//
//func TestGetExistingCart(t *testing.T)  {
//	setup()
//
//	method := "GET"
//	urlStr := "/cart/321"
//
//	req, err := http.NewRequest(method, urlStr, nil)
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//
//	muxRouter.ServeHTTP(respRec, req)
//
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
//	}
//
//}
//
//func TestAddItemToExistingCart(t *testing.T) {
//	setup()
//
//	method := "POST"
//	urlStr := "/cart/423"
//
//	postBody := map[string]interface{}{
//		"prodId": 2,
//		"qty": 23,
//		"price":234,
//	}
//
//	body, _ := json.Marshal(postBody)
//	req, err := http.NewRequest(method, urlStr, bytes.NewReader(body))
//	if err != nil {
//		t.Fatal("Creating " + method + " " + urlStr + " request failed!")
//	}
//	muxRouter.ServeHTTP(respRec, req)
//	t.Log("Returned Body: ", respRec.Body)
//
//	if respRec.Code != http.StatusOK {
//		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusOK)
//	}
//}