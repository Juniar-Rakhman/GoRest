package router

import (
	"jrakhman/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//Product Routes
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GetProductList",
		"GET",
		"/product",
		handlers.GetProductAll,
	},
	Route{
		"GetProductById",
		"GET",
		"/product/{productId}",
		handlers.GetProductById,
	},
	Route{
		"AddProduct",
		"POST",
		"/product",
		handlers.AddProduct,
	},
	Route{
		"DelProduct",
		"DELETE",
		"/product/{id}",
		handlers.DelProduct,
	},
	//Cart Routes

}
