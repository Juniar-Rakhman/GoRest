package repo

import (
	"fmt"
	"jrakhman/model"
	_ "github.com/lib/pq"
	"database/sql"
	//"time"
)

//TODO: Replace this package with proper ORM ?

const (
	DRIVER = "postgres"
	URL = "postgresql://qfheqofagbnxly:cv0Fe77903nXAUr6GFm3reOWHL@ec2-23-21-50-120.compute-1.amazonaws.com/df2t58f1j6163a"
)

var currentId int
var productList model.Products

// create some seed data
//func init() {
//	productList = model.Products{}
//	CreateProduct(model.Product{Name: "Sandal Jepit", Size: 32, Color:"red", Price:100})
//	CreateProduct(model.Product{Name: "Celana Panjang", Size: 30, Color:"blue", Price:500})
//	CreateProduct(model.Product{Name: "Kemeja", Size: 30, Color:"green", Price:500})
//}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FindAll() model.Products {

	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_product")
	checkErr(err)

	response := model.Products{}

	for rows.Next() {
		p := model.Product{}

		err = rows.Scan(&p.Id, &p.Name, &p.Size, &p.Color, &p.Price)
		checkErr(err)

		response.Products = append(response.Products, p)
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
