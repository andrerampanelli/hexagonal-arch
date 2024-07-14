package db

import (
	"database/sql"

	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	"github.com/andrerampanelli/hexagonal-arch/application/interfaces"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
}

func (p *ProductDb) Get(id string) (interfaces.ProductInterface, error) {
	var product domain.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	var rows int
	err := p.db.QueryRow("SELECT COUNT(id) FROM products WHERE id = ?", product.GetId()).Scan(&rows)
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return p.create(product)
	}
	return p.update(product)
}

func (p *ProductDb) create(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetId(),
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) Delete(product interfaces.ProductInterface) error {
	stmt, err := p.db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.GetId())
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductDb) List() ([]interfaces.ProductInterface, error) {
	var products []interfaces.ProductInterface

	rows, err := p.db.Query("SELECT id, name, price, status FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
