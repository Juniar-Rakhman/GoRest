package router

import "github.com/jrakhman/model"

var routes = model.Routes {

	model.Route{
		"ProductShow",
		"GET",
		"/products/{stock}",
		productController.GetStockByText,
	},
}

