package item

import (
	"context"

	"github.com/LibenHailu/2FInventory/internal/constant/model"
	"github.com/LibenHailu/2FInventory/internal/storage/persistance"
)

// business logic for the domain  item
type Usecase interface {
	InsertItem(ctx context.Context, item *model.Item) (*model.Item, error)
}

type Service struct {
	itemPersist persistance.ItemPersistance
}

// initializing our service
func Initialize(itemPersist persistance.ItemPersistance) Usecase {
	return &Service{
		itemPersist: itemPersist,
	}
}
