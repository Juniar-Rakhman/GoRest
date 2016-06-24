package handlers

import (
	"fmt"
	"github.com/jrakhman/model"
)

var currentId int

var productList model.ProductList

// create some seed data
func init() {
	RepoCreateProduct(model.Product{Name: "Sandal Jepit", Size: 32, Categories:[]string{"fashion", "footwear"}, AvailableStock:21, Price:100})
	RepoCreateProduct(model.Product{Name: "Celana Panjang", Size: 30, Categories:[]string{"fashion", "pants"}, AvailableStock:30, Price:500})
}

func RepoFindProduct(id int) model.Product {
	for _, p := range productList {
		if p.Id == id {
			return p
		}
	}
	// return empty product if not found
	return model.Product{}
}

func RepoCreateProduct(p model.Product) model.Product {
	currentId += 1
	p.Id = currentId
	productList = append(productList, p)
	return p
}

func RepoDeleteProduct(id int) error {
	for i, p := range productList {
		if p.Id == id {
			productList = append(productList[:i], productList[i + 1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
