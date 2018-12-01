package queries

// GetCustomersQuery return customers SQL query
func GetCustomersQuery() string {
	return "SELECT customers_id,customers_gender,customers_firstname,customers_lastname,customers_email_address,customers_telephone FROM customers"
}
