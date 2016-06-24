package router

import (
	"github.com/jrakhman/handlers"
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
		"GetProductByCategory",
		"GET",
		"/categories/{cat:[0-9]+}",
		handlers.GetProductByCategory,
	},
	Route{
		"GetProductByPrice",
		"GET",
		"/price/{max}/{min}",
		handlers.GetProductByPrice,
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
}
