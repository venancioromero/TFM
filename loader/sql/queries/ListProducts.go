package queries

// GetProductsQuery return customers SQL query
func GetProductsQuery() string {
	return `SELECT  products.products_id,
					products_to_categories.categories_id,
					products.products_image,
					products.products_status,
					products.products_ordered

			FROM 	products,products_to_categories 

			WHERE 	products.products_id=products_to_categories.products_id`
}
