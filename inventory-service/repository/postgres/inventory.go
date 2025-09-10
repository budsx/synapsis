package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/budsx/synapsis/inventory-service/entity"
	"github.com/budsx/synapsis/inventory-service/repository/interfaces"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dbHost string, dbPort int, dbUser string, dbPassword string, dbName string, dbDriver string) (interfaces.InventoryDBReadWriter, error) {
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

func (r *postgresRepository) CheckStock(ctx context.Context, productID int64) (int64, error) {
	query := `SELECT stock FROM products WHERE id = $1`

	var stock int64
	err := r.db.QueryRowContext(ctx, query, productID).Scan(&stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return stock, nil
}

func (r *postgresRepository) ReserveStock(ctx context.Context, productID int64, quantity int64) error {
	querySelect := `SELECT id, stock, version FROM products WHERE id = $1`
	var version, stock int64
	err := r.db.QueryRowContext(ctx, querySelect, productID).Scan(&version)
	if err != nil {
		return err
	}
	if stock < quantity {
		return fmt.Errorf("stock is not enough")
	}

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryUpdate := `UPDATE products SET stock = stock - $1 WHERE id = $2 AND version = $3`
	result, err := tx.ExecContext(ctx, queryUpdate, quantity, productID, version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("stock is not enough")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) ReleaseStock(ctx context.Context, productID int64, quantity int64) error {
	querySelect := `SELECT id, stock, version FROM products WHERE id = $1`
	var version int64
	err := r.db.QueryRowContext(ctx, querySelect, productID).Scan(&version)
	if err != nil {
		return err
	}
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryUpdate := `UPDATE products SET stock = stock + $1 WHERE id = $2 AND version = $3`
	result, err := tx.ExecContext(ctx, queryUpdate, quantity, productID, version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("release stock failed will re retry immediately")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *postgresRepository) GetProductByID(ctx context.Context, productID int64) (*entity.Product, error) {
	query := `SELECT id, name, stock, created_at, updated_at FROM products WHERE id = $1`

	var product entity.Product
	err := r.db.QueryRowContext(ctx, query, productID).Scan(
		&product.ID,
		&product.Name,
		&product.Stock,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *postgresRepository) Close() error {
	return r.db.Close()
}
