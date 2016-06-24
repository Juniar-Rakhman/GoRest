package model

import "net/http"

// product container
type ProductList struct {
	ProductList []Product `json:"Product"`
}

// product
type Product struct {
	Name string
	Size string
	Color string
	Prize string
}

// routes
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route