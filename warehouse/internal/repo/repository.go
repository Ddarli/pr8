package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Ddarli/app/warehouse/config"
	"github.com/Ddarli/app/warehouse/pkg"
	"github.com/Ddarli/utils/models"
	_ "github.com/lib/pq"
	"log"
)

type (
	WarehouseRepo struct {
		db        *sql.DB
		converter pkg.Converter
	}
)

func NewWarehouseRepo() Repository {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &WarehouseRepo{
		db:        db,
		converter: &pkg.ProductConverter{},
	}
}

func (r *WarehouseRepo) GetAll(ctx context.Context) ([]*models.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, quantity FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []pkg.Product
	for rows.Next() {
		var product pkg.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}

	return r.converter.ToProto(products), err
}

func (r *WarehouseRepo) CheckQuantity(ctx context.Context, id int, quantity int) (bool, error) {
	var q int
	err := r.db.QueryRow("SELECT quantity FROM products WHERE id=$1", id).Scan(&q)
	if err != nil {
		return false, err
	}

	if q >= quantity {
		return true, nil
	}

	return false, nil
}
