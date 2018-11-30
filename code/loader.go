package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type row struct {
	OrderID    int
	CustomerID int
	ProductID  int
	ImageURL   string
	ProductURL string
}

func (r row) String() string {
	return fmt.Sprint("OrderId:" + strconv.Itoa(r.OrderID) + "  " +
		"CustomerID:" + strconv.Itoa(r.CustomerID) + "  " +
		"ProductID:" + strconv.Itoa(r.ProductID) + "  " +
		"ImageURL:" + r.ImageURL + "  " +
		"ProductURL:" + r.ProductURL)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	fmt.Println("Loader Running...")
	db, err := sql.Open("mysql", "root:1234@/panel555_RS")
	checkErr(err)
	defer db.Close()

	// Execute the query
	results, err := db.Query(`SELECT 	orders.orders_id ,
									   	orders.customers_id,
									   	orders_products.products_id, 
									   	CONCAT("/images/",products.products_image) as image_url,
										CONCAT("https://www.panel555.com/product_info.php?products_id=",orders_products.products_id) as url_product 
   								FROM
	 									orders,products,orders_products
   								WHERE 
										orders.orders_id         = orders_products.orders_id AND
										products.products_id     = orders_products.products_id AND 
										products.products_status = 1 AND
										orders.date_purchased    > "2015-00-00 00:00:00"
								ORDER BY 
										orders_id DESC`)

	checkErr(err)

	for results.Next() {

		var r row
		err = results.Scan(&r.OrderID, &r.CustomerID, &r.ProductID, &r.ImageURL, &r.ProductURL)
		checkErr(err)
		fmt.Println(r)
	}

}
