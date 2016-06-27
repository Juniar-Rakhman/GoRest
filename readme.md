Product Shopping cart supporting: add, remove item to cart, support for discount coupon and total purchase amount

Tested environment:
- OS: MacOSX
- IDE: Intellij with Go plugins
- Go version: 1.6

To run the app : 
- run 'go build' from root directory 
- or run executable binary in bin folder

To clean & fix dependencies:
- delete folder from src/github.com/*
- run 'go get' from root directory

To run unit test : 
- uncomment the test cases that you want to run in test/main_test.go
- run 'go test -v' from test directory

ROUTES:
From src/jrakhman/router/routes.go

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
