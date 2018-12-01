package queries

// GetOrdersQuery will returns query that retrieve orders data
func GetOrdersQuery() string {
	return `SELECT 	orders.orders_id ,
					orders.customers_id,
					orders_products.products_id,
					orders.date_purchased
   			FROM
	 				orders,orders_products
   			WHERE 
					orders.orders_id         = orders_products.orders_id AND
					orders.date_purchased    > "2015-00-00 00:00:00"
					ORDER BY orders_id DESC;
			`
}
