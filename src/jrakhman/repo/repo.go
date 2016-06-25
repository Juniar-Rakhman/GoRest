package repo

import (
	"fmt"
	"jrakhman/model"
)

//TODO: Replace this package with proper ORM ?

var currentId int

var productList model.Products

// create some seed data
func init() {
	productList = model.Products{}
	CreateProduct(model.Product{Name: "Sandal Jepit", Size: 32, Color:"red", Price:100})
	CreateProduct(model.Product{Name: "Celana Panjang", Size: 30, Color:"blue", Price:500})
	CreateProduct(model.Product{Name: "Kemeja", Size: 30, Color:"green", Price:500})
}

func FindAll() model.Products {

	response := model.Products{}

	for _, prod := range productList.Products {
		response.Products = append(response.Products, prod)
	}

	return response
}

func FindProduct(id int) model.Product {
	for _, p := range productList.Products {
		if p.Id == id {
			return p
		}
	}
	// return empty product if not found
	return model.Product{}
}

func CreateProduct(p model.Product) model.Product {
	currentId += 1
	p.Id = currentId
	productList.Products = append(productList.Products, p)
	return p
}

func DeleteProduct(id int) error {
	//for i, p := range productList.Products {
	//	if p.Id == id {
	//		productList = append(productList[:i], productList[i + 1:]...)
	//		return nil
	//	}
	//}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
