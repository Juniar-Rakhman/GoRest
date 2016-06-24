package model

type Product struct {
	Id         int		`json:"id"`
	Name       string	`json:"name"`
	Size       int		`json:"size"`
	Categories []string	`json:"category"`
	AvailableStock  int	`json:"available"`
	Price      int		`json:"price"`
}

type ProductList []Product
