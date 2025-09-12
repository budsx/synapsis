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

func (r *postgresRepository) CreateOrder(ctx context.Context, req *entity.CreateOrderRequest) (*entity.CreateOrderResponse, error) {
	return nil, nil
}

func (r *postgresRepository) GetOrderByID(ctx context.Context, req *entity.GetOrderByIDRequest) (*entity.Order, error) {
	return nil, nil
}

func (r *postgresRepository) Close() error {
	return r.db.Close()
}
