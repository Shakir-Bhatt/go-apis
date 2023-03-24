package controllers

import (
	"fmt"

	"simpleapi_go/database"

	_ "github.com/go-sql-driver/mysql"

	"simpleapi_go/models"
)

// func GetProducts() []models.Product {
func GetProducts() []models.ProductWithImgae {

	// GORM
	//

	// call NewMysqlConnection to get database instance
	db := database.NewMysqlConnection()

	var fields = []string{"id", "title"}
	returnData := []models.ProductWithImgae{}

	fmt.Println(fields)
	results, err := db.Query("SELECT * FROM products")

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil

	}

	// products := []models.Product{}

	for results.Next() {

		// var prod models.Product

		// err = results.Scan(&prod)

		// var obj models.ProductWithImgae

		// obj.ID = prod.ID
		// obj.ImageUrl = prod.Image.Src
		// obj.Title = prod.Title
		// obj.Vendor = prod.Vendor

		// returnData = append(returnData, obj)

		// if err != nil {
		// 	panic(err.Error())
		// }

		// products = append(products, prod)

		// removeEmptyFields(products)

		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)

		// scan the row into a map
		row := make(map[string]interface{})
		err = results.Scan(row)
		if err != nil {
			panic(err.Error())
		}

		// create a new ProductWithImage struct
		obj := models.ProductWithImgae{
			ID:       row["id"].(int64),
			Title:    row["title"].(string),
			ImageUrl: row["image_url"].(string),
			Vendor:   row["vendor"].(string),
		}

		// add the struct to the return data slice
		returnData = append(returnData, obj)
	}

	return returnData

}

func GetProduct(id string) *models.Product {

	db := database.NewMysqlConnection()

	prod := &models.Product{}

	results, err := db.Query("SELECT * FROM products where id=?", id)

	if err != nil {

		fmt.Println("Err", err.Error())

		return nil
	}

	if results.Next() {

		// err = results.Scan(&prod.Code, &prod.Name, &prod.Qty, &prod.LastUpdated)

		// if err != nil {
		// 	return nil
		// }
	} else {

		return nil
	}

	return prod

}

// func AddProduct(product models.Product) {

// 	db := database.NewMysqlConnection()

// 	insert, err := db.Query(
// 		"INSERT INTO products (code,name,qty,last_updated) VALUES (?,?,?, now())",
// 		product.Code, product.Name, product.Qty)

// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer insert.Close()

// }

// func AddProduct() int {

// 	db := database.NewMysqlConnection()

// 	var prod models.Product

// 	if err := payload.BindJSON(&prod); err != nil {
// 		return 500
// 	}

// 	_, err := db.Query(
// 		"INSERT INTO product (code,name,qty,last_updated) VALUES (?,?,?, now())",
// 		product.Code, product.Name, product.Qty)

// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// defer insert.Close()

// }
