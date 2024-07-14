package db_test

import (
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"log"
	"testing"

	"github.com/andrerampanelli/hexagonal-arch/adapters/db"
	"github.com/andrerampanelli/hexagonal-arch/application/domain"
	"github.com/stretchr/testify/require"
)

var Conn *sql.DB

func setUp() {
	Conn, _ = sql.Open("sqlite3", "file::memory:?cache=shared")
	createTable(Conn)
	insertProduct(Conn)
}

func createTable(conn *sql.DB) {
	table := `CREATE TABLE products (
		"id"     string primary key,
		"name"   string,
		"price"  float,
		"status" string
	)`

	stmt, err := conn.Prepare(table)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func insertProduct(conn *sql.DB) {
	insertProduct := `INSERT INTO products VALUES ("1", "Product 1", 0, "disabled")`

	stmt, err := conn.Prepare(insertProduct)
	if err != nil {
		panic(err)
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Conn.Close()
	productDb := db.NewProductDb(Conn)

	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "1", product.GetId())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

	product, err = productDb.Get("ABC")
	require.NotNil(t, err)
	require.Nil(t, product)
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Conn.Close()
	productDb := db.NewProductDb(Conn)

	product := &domain.Product{
		Id:     "2",
		Name:   "Product 2",
		Price:  10.99,
		Status: "enabled",
	}

	// Test create
	createdProduct, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetId(), createdProduct.GetId())
	require.Equal(t, product.GetName(), createdProduct.GetName())
	require.Equal(t, product.GetPrice(), createdProduct.GetPrice())
	require.Equal(t, product.GetStatus(), createdProduct.GetStatus())

	// Test update
	product.Name = "Updated Product 2"
	product.Price = 19.99
	updatedProduct, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetId(), updatedProduct.GetId())
	require.Equal(t, product.GetName(), updatedProduct.GetName())
	require.Equal(t, product.GetPrice(), updatedProduct.GetPrice())
	require.Equal(t, product.GetStatus(), updatedProduct.GetStatus())
}
