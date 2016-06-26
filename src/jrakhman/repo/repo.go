package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"jrakhman/model"
	//"time"
)

/*
	This package responsible for handling DB CRUD
*/

//TODO: Replace this package with proper ORM ? hmm...slower performance?

const (
	DRIVER = "postgres"
	URL    = "postgresql://qfheqofagbnxly:cv0Fe77903nXAUr6GFm3reOWHL@ec2-23-21-50-120.compute-1.amazonaws.com/df2t58f1j6163a"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FindAll() model.Products {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close() //always close db at callback. Since we use free DB, connection is limited.

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

func DeleteProduct(id int) string {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("delete from tbl_product where id=$1")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	response := "deleted product id = " + string(id)

	return response
}

//--- Carts ----//

func AddItemToNewCart(user int, ci model.CartItem) model.Cart {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	var lastCartItemId int
	var lastCartId int

	c := model.Cart{}
	c.UserId = user
	c.Total = ci.Price
	c.Discount = 0
	c.Paid = false

	//insert cart
	err = db.QueryRow("INSERT INTO tbl_cart_payment(user_id, total, discount, paid) VALUES($1,$2,$3,$4) returning id;", c.UserId, c.Total, c.Discount, c.Paid).Scan(&lastCartId)
	checkErr(err)

	fmt.Println("last inserted cart id=", lastCartId)
	c.Id = lastCartId

	c.CartItems = append(c.CartItems, ci)

	//insert cart item
	err = db.QueryRow("INSERT INTO tbl_cart_items(cart_id, product_id, quantity, price) VALUES($1,$2,$3,$4) returning id;", lastCartId, ci.ProdId, ci.Quantity, ci.Price).Scan(&lastCartItemId)
	checkErr(err)

	fmt.Println("last inserted cart item id =", lastCartItemId)

	return c
}

func AddItemToExistingCart(userId int, ci model.CartItem) model.Cart {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	var lastCartItemId int

	c := FindCartByUser(userId)

	//insert cart item
	err = db.QueryRow("INSERT INTO tbl_cart_items(cart_id, product_id, quantity, price) VALUES($1,$2,$3,$4) returning id;", c.Id, ci.ProdId, ci.Quantity, ci.Price).Scan(&lastCartItemId)
	checkErr(err)

	println(lastCartItemId)

	fmt.Println("last inserted cart item id =", lastCartItemId)

	c.CartItems = append(c.CartItems, ci)

	CalculateTotalCost(&c)

	return c
}

func FindCartByUser(userId int) model.Cart {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	c := model.Cart{}
	cRows, err := db.Query("SELECT * FROM tbl_cart_payment where user_id=$1", userId)
	checkErr(err)

	for cRows.Next() {
		//should be only 1
		err = cRows.Scan(&c.Id, &c.Total, &c.Discount, &c.Paid, &c.UserId)
		checkErr(err)
	}

	ciRows, err := db.Query("SELECT * FROM tbl_cart_items where cart_id=$1", c.Id)
	checkErr(err)

	for ciRows.Next() {
		var cart int
		ci := model.CartItem{}
		err = ciRows.Scan(&ci.Id, &cart, &ci.ProdId, &ci.Quantity, &ci.Price)
		checkErr(err)
		c.CartItems = append(c.CartItems, ci)
	}

	return c
}

//func DeleteCartItem(userId int) model.Cart {
//
//}
//
//func SetCartToPaid(userId int) model.Cart {
//
//}
//
//func AddPromoCode(userId int, discount int) model.Cart {
//
//}

func CalculateTotalCost(c *model.Cart) {
	db, err := sql.Open(DRIVER, URL)
	checkErr(err)
	defer db.Close()

	var totalCost int

	for _, ci := range c.CartItems {
		totalCost += ci.Price
	}

	c.Total = totalCost

	fmt.Println("# Updating")
	stmt, err := db.Prepare("update tbl_cart_payment set total=$1 where id=$2")
	checkErr(err)

	res, err := stmt.Exec(c.Total, c.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")
}
