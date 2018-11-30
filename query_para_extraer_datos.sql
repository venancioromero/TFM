SELECT orders.orders_id ,
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
        ORDER BY orders_id DESC;