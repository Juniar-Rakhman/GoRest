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
	Route{	//will send cartItem json
		"AddItemToNewCart",
		"POST",
		"/cart/new/{userId}",
		handlers.AddItemToNewCart,
	},
	Route{
		"GetExistingCart",
		"GET",
		"/cart/{userId}",
		handlers.GetExistingCartByUser,
	},
	Route{
		"AddItemToExistingCart",
		"POST",
		"/cart/{userId}",
		handlers.AddItemToExistingCart,
	},
	Route{
		"DeleteItemFromCart",
		"DELETE",
		"/cart/item/{itemId}",
		handlers.DeleteCartItem,
	},
	//Payment Routes
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
		handlers.AddDiscount,
	},
}
