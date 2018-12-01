package models

import (
	"database/sql"
	"strings"
)

// Customer Model
type Customer struct {
	ID        sql.NullInt64
	Gender    sql.NullString
	Firstname sql.NullString
	Lastname  sql.NullString
	Email     sql.NullString
	Telephone sql.NullString
}

// Node return Neo4j Template
func (c Customer) Node() string {
	return "CREATE (n:customer {customerID: {ID}, gender: {gender},fullname:{fullname},email:{email},phone:{phone}})"
}

//GetProperties will return map with all properties of customer
func (c Customer) GetProperties() map[string]interface{} {
	return map[string]interface{}{"ID": c.ID.Int64, "gender": strings.ToUpper(c.Gender.String), "fullname": c.Firstname.String, "email": strings.ToLower(c.Email.String), "phone": c.Telephone.String}
}
