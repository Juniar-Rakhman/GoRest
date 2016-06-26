package model

type CartItem struct {
	ProdId  	int `json:"prodId"`
	Quantity	int `json:"qty"`
	Price		int `json:"price"`
}

type Cart struct {
	Id  		int `json:"id"`
	UserId 		int `json:"user"`
	Total		int `json:"total"`
	Discount 	int `json:"discount"`
	Paid 		bool `json:"paid"`
	CartItems   []CartItem  `json:"cartItems"`
}