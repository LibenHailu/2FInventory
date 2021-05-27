package persistance

import (
	"context"
	"errors"
	"log"

	"github.com/LibenHailu/2FInventory/internal/constant/model"
	"github.com/LibenHailu/2FInventory/internal/constant/query"
	"github.com/jackc/pgx"
)

// methods for item database table
type ItemPersistance interface {
	InsertItem(ctx context.Context, item *model.Item) (*model.Item, error)
}

type itemPersistance struct {
	conn *pgx.Conn
}

func ItemInit(conn *pgx.Conn) ItemPersistance {
	return &itemPersistance{
		conn: conn,
	}
}

func (it *itemPersistance) InsertItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	newItem, err := it.conn.Exec(ctx, query.QueryInsertItem, item.Name, item.Price, item.Code, item.Barcode, item.ExpireDate, item.WarningDate, item.Quantity, item.RegisteredBy)

	if err != nil {
		log.Fatalf("some thing wrong: %v", err)
	}
	if newItem.RowsAffected() > 0 {

		return item, nil
	}
	return nil, errors.New("Some thing went wrong")
}
