package model

type CartItem struct {
	ProdId  	int `json:"productId"`
	Quantity	int `json:"productQty"`
}

type Cart struct {
	Id  		int `json:"id"`
	UserId 		int `json:"user"`
	CartItems   []CartItem `json:"cartItems"`
}

type CartPayment struct {
	Total		int `json:"total"`
	Discount 	int `json:"discount"`
	Paid 		bool `json:"paid"`
	Cart		Cart `json:"cart"`
}