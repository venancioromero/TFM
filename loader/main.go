package main

import (
	"database/sql"
	"fmt"

	"./config"
	"./models"

	queries "./sql/queries"
	_ "github.com/go-sql-driver/mysql"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

var conf config.Config
var sqlConn *sql.DB
var neoConn bolt.Conn

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	conf = config.New()

	// mySQL
	db, err := sql.Open("mysql", conf.SQLURL)
	checkErr(err)
	sqlConn = db

	// Neo4j
	neo, err := bolt.NewDriver().OpenNeo(conf.Neo4JURL)
	checkErr(err)
	neoConn = neo

}

func insertCustomers() {
	fmt.Println("\tInserting customers...")
	var c models.Customer
	query := queries.GetCustomersQuery()
	results, err := sqlConn.Query(query)
	checkErr(err)

	for results.Next() {
		err = results.Scan(&c.ID, &c.Gender, &c.Firstname, &c.Lastname, &c.Email, &c.Telephone)
		checkErr(err)
		_, err := neoConn.ExecNeo(c.Node(), c.GetProperties())
		checkErr(err)
		fmt.Print()
	}

}
func insertProducts() {
	fmt.Println("\tInserting Products...")
	var p models.Product
	query := queries.GetProductsQuery()
	results, err := sqlConn.Query(query)
	checkErr(err)

	for results.Next() {
		err = results.Scan(&p.ID, &p.Category, &p.Image, &p.Status, &p.ProductsOrdered)
		checkErr(err)
		_, err := neoConn.ExecNeo(p.Node(), p.GetProperties())
		checkErr(err)
	}
}

func insertOrders() {
	fmt.Println("\tInserting Orders...")
	var o models.Order
	query := queries.GetOrdersQuery()
	results, err := sqlConn.Query(query)
	checkErr(err)

	for results.Next() {
		err = results.Scan(&o.OrderID, &o.CustomerID, &o.ProductID, &o.DatePurchase)
		checkErr(err)
		_, err := neoConn.ExecNeo(o.Edge(), o.GetProperties())
		checkErr(err)
	}
}

func main() {
	fmt.Println("Loader Running...")
	insertCustomers()
	insertProducts()
	insertOrders()
	defer sqlConn.Close()
	defer neoConn.Close()
}
