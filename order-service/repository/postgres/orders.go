package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/budsx/synapsis/order-service/entity"
	"github.com/budsx/synapsis/order-service/repository/interfaces"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dbHost string, dbPort int, dbUser string, dbPassword string, dbName string, dbDriver string) (interfaces.OrderDBReadWriter, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &postgresRepository{
		db: db,
	}, nil
}

func (r *postgresRepository) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (int64, error) {
	query := `INSERT INTO orders (product_id, quantity, status, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	result, err := r.db.ExecContext(ctx, query, req.ProductID, req.Quantity, req.Status)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (r *postgresRepository) UpdateOrderStatus(ctx context.Context, req *entity.UpdateOrderStatusRequest) error {
	query := `UPDATE orders SET status = $1 WHERE order_id = $2`
	_, err := r.db.ExecContext(ctx, query, req.Status, req.OrderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepository) Close() error {
	return r.db.Close()
}
