package model

type CartItem struct {
	ProdId  	int `json:"prodId"`
	Quantity	int `json:"qty"`
	Price		int `json:"price"`
}

type Cart struct {	//we should have 1 cart = 1 user
	Id  		int `json:"id"`
	UserId 		int `json:"user"`
	Total		int `json:"total"`
	Discount 	int `json:"discount"`
	Paid 		bool `json:"paid"`
	CartItems   []CartItem  `json:"cartItems"`
}