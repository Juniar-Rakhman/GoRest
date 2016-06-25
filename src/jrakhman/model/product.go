package model

type Product struct {
	Id         int		`json:"id"`
	Name       string	`json:"name"`
	Size       int		`json:"size"`
	Color  		string	`json:"color"`
	Price      int		`json:"price"`
}

type Products struct {
	Products []Product `json:"Products"`
}
