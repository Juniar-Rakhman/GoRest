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
		"/product/{productId}",
		handlers.DeleteProduct,
	},
	//Cart Routes
	Route{
		"GetExistingCart",
		"GET",
		"/cart/{cartId}",
		handlers.GetExistingCart,
	},
	Route{
		"AddItemToNewCart",
		"POST",
		"/cart/new",
		handlers.AddItemToNewCart,
	},
	Route{
		"AddItemToExistingCart",
		"POST",
		"/cart/{cartId}",
		handlers.AddItemToExistingCart,
	},
	Route{
		"DeleteItemFromCart",
		"DELETE",
		"/cart/{cartId}/{itemId}",
		handlers.DeleteItemFromCart,
	},
	//Payment Routes
	Route{
		"GetPaymentDetails",
		"GET",
		"/payment/{cartId}",
		handlers.GetPaymentDetails,
	},
	Route{
		"SetCartToPaid",
		"POST",
		"/payment/{cartId}",
		handlers.SetCartToPaid,
	},
	Route{
		"AddDiscountCode",
		"POST",
		"/payment/{cartId}/{discountCode}",
		handlers.AddDiscountCode,
	},
}
