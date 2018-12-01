package models

import (
	"database/sql"
)

// Product Model
type Product struct {
	ID              sql.NullInt64
	Category        sql.NullInt64
	Image           sql.NullString
	Status          sql.NullBool
	ProductsOrdered sql.NullInt64
}

// Node return Neo4j Template
func (p Product) Node() string {
	return "CREATE (n:product {productID: {ID}, category: {category},image:{image},status:{status},ordered:{ordered}})"
}

//GetProperties will return map with all properties of customer
func (p Product) GetProperties() map[string]interface{} {
	return map[string]interface{}{"ID": p.ID.Int64, "category": p.Category.Int64, "image": p.Image.String, "status": p.Status.Bool, "ordered": p.ProductsOrdered.Int64}
}
