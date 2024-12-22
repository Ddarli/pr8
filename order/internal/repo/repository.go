package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Ddarli/app/order/config"
	"github.com/Ddarli/app/order/pkg/models"
	_ "github.com/lib/pq"
	"log"
)

type (
	OrderRepo struct {
		db *sql.DB
	}
)

func New() Repository {
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

	return &OrderRepo{
		db: db,
	}
}

func (r *OrderRepo) SaveOrder(ctx context.Context, order *models.Order) (bool, error) {
	query := "INSERT INTO orders (customer_name, order_date, total_amount) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, order.Customer, order.Date, order.Total)
	if err != nil {
		return false, err
	}
	return true, nil
}
