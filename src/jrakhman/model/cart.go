package model

type CartItems struct {
	Id  	int `json:"id"`
	UserId 	int `json:"user"`
	CartNo	int `json:"cart"`
	ProdId  int `json:"productId"`
	Qty     int `json:"productQty"`
	Disc	int `json:"discount"`
}