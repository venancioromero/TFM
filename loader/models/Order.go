package models

import (
	"database/sql"
)

// Order Model
type Order struct {
	OrderID      sql.NullInt64
	CustomerID   sql.NullInt64
	ProductID    sql.NullInt64
	DatePurchase sql.NullString
}

// Edge return Neo4j Template
func (o Order) Edge() string {
	return "MATCH (cust:customer {customerID:{customerID}}),(prod:product {productID:{productID}}) CREATE (cust)-[:BOUGHT{datePurchase:{datePurchase}}]->(prod)"
}

//GetProperties will return map with all properties of customer
func (o Order) GetProperties() map[string]interface{} {
	return map[string]interface{}{"customerID": o.CustomerID.Int64, "productID": o.ProductID.Int64, "datePurchase": o.DatePurchase.String}
}
