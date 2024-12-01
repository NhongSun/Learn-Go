package main

func createProduct(product *Product) error {
	// Exec executes a query without returning any rows.
	_, err := db.Exec("INSERT INTO products(name, price) VALUES ($1, $2)", product.Name, product.Price)

	return err
}

func getProduct(id int) (Product, error) {
	var p Product

	// QueryRow returns a row
	// QueryRow executes a query that is expected to return at most one row.
	// QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called.
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id=$1", id)

	// Scan copies the columns in the current row into the values pointed at by dest.
	// The number of values in dest must be the same as the number of columns in Rows.
	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product

	row := db.QueryRow("UPDATE products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price", product.Name, product.Price, id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func deleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", id)

	return err
}

func getProducts() ([]Product, error) {
	// Query returns a Rows object
	// Query executes a query that returns rows
	rows, err := db.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func getProductsWithSuppliers() ([]ProductWithSupplier, error) {
	// column name must be ht same as the struct field name
	query := `
		SELECT 
			p.id AS product_id,
			p.name AS product_name,
			p.price,
			s.name AS supplier_name
		FROM
			products p INNER JOIN suppliers s
			ON p.supplier_id = s.id
		ORDER BY product_id
	`

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	var products []ProductWithSupplier

	for rows.Next() {
		var p ProductWithSupplier
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.SupplierName)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
