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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FindAll() model.Products {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()    //always close db at callback. Since we use free DB, connection is limited.

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
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_product where id=$1", id)
	checkErr(err)

	p := model.Product{}

	for rows.Next() {
		//should be only 1
		err = rows.Scan(&p.Id, &p.Name, &p.Size, &p.Color, &p.Price)
		checkErr(err)
	}

	return p
}

func CreateProduct(p model.Product) model.Product {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	var lastInsertId int

	err = db.QueryRow("INSERT INTO tbl_product(name,size,color,price) VALUES($1,$2,$3,$4) returning id;", p.Name, p.Size, p.Color, p.Price).Scan(&lastInsertId)
	checkErr(err)

	fmt.Println("last inserted id =", lastInsertId)

	p.Id = lastInsertId

	return p
}

func DeleteProduct(id int) error {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	_, err = stmt.Exec(id)

	fmt.Println("deleted product id = ", id)

	return err
}
