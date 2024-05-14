package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/models"
	repository "github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/postgres"
)

type WarehouseRepos struct {
	db *repository.DB
}

func NewWarehouseRepos(db *repository.DB) *WarehouseRepos {
	return &WarehouseRepos{db: db}
}

func (r *WarehouseRepos) GetItems(ctx context.Context, id int64) ([]models.ItemModel, error) {
	var items []models.ItemModel

	query := `
        SELECT i.name, i.size, i.unique_code, iwq.quantity
        FROM items i
        JOIN item_warehouse_quantity iwq ON i.id = iwq.item_id
        WHERE iwq.warehouse_id = $1
    `

	rows, err := r.db.DB.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.ItemModel
		if err := rows.Scan(&item.Name, &item.Size, &item.UniqueCode, &item.Quantity); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *WarehouseRepos) ReservedItems(ctx context.Context, reservations []models.ReservationModel) (bool, error) {
	insertReservedItemQuery := `
  		INSERT INTO reserved_items (item_warehouse_id, quantity)
  		VALUES ($1, $2)
 	`
	updateItemWarehouseQuantity := `
  		UPDATE item_warehouse_quantity
		SET quantity = quantity - $1
  		WHERE id = $2;
 	`

	err := r.withTransaction(ctx, func(tx pgx.Tx) error {
		for _, res := range reservations {
			itemWarehouseID, err := r.itemWarehouseID(ctx, tx, res.UniqueCode, res.WarehouseId)
			if err != nil {
				return err
			}

			_, err = tx.Exec(ctx, insertReservedItemQuery, itemWarehouseID, res.Quantity)
			if err != nil {
				return err
			}

			_, err = tx.Exec(ctx, updateItemWarehouseQuantity, res.Quantity, itemWarehouseID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err == nil, err
}

func (r *WarehouseRepos) ReleaseReserve(ctx context.Context, reservations []models.ReservationModel) (bool, error) {
	deleteReservedItems := `
  		DELETE FROM reserved_items
		WHERE ctid = (
			SELECT ctid FROM reserved_items
			WHERE item_warehouse_id = $1 AND quantity = $2
			LIMIT 1
		)
		RETURNING id;
 	`

	err := r.withTransaction(ctx, func(tx pgx.Tx) error {
		for _, res := range reservations {
			itemWarehouseID, err := r.itemWarehouseID(ctx, tx, res.UniqueCode, res.WarehouseId)
			if err != nil {
				return err
			}

			var deletedID int
			err = tx.QueryRow(ctx, deleteReservedItems, itemWarehouseID, res.Quantity).Scan(&deletedID)
			if err != nil {
				return errors.Wrap(err, "no row deleted")
			}

			if deletedID == 0 {
				return errors.Wrap(err, "no row deleted")
			}
		}
		return nil
	})

	return err == nil, err
}

func (r *WarehouseRepos) itemWarehouseID(ctx context.Context, tx pgx.Tx, uniqueCode, warehouseId int) (int64, error) {
	selectItemWarehouseID := `
  		SELECT iwq.id
  		FROM item_warehouse_quantity iwq
  		JOIN items i ON iwq.item_id = i.id
  		JOIN warehouse w ON iwq.warehouse_id = w.id
  		WHERE i.unique_code = $1 AND w.id = $2
 	`

	var itemWarehouseID int64

	err := tx.QueryRow(ctx, selectItemWarehouseID, uniqueCode, warehouseId).Scan(&itemWarehouseID)
	if errors.Is(err, pgx.ErrNoRows) {
		return itemWarehouseID, errors.Wrap(err, "invalid item unique code or warehouse id")
	} else if err != nil {
		return itemWarehouseID, err
	}

	return itemWarehouseID, nil
}

func (r *WarehouseRepos) withTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := r.db.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				err = errors.Wrap(errTx, "Rollback FAILED")
			}

			log.Println("ROLLBACK IS OK.")
		} else {
			if errTx := tx.Commit(ctx); errTx != nil {
				err = errors.Wrap(errTx, "Commit FAILED")
			}

			log.Println("COMMIT IS OK")
		}
	}()

	err = fn(tx)

	return err
}
