package item

import (
	"context"
	"log"

	"github.com/LibenHailu/2FInventory/internal/constant/model"
)

// Inserts a new Item to the db
func (s *Service) InsertItem(ctx context.Context, item *model.Item) (*model.Item, error) {
	item, err := s.itemPersist.InsertItem(ctx, item)

	if err != nil {
		log.Printf("could not inset item: %v", err)
		return nil, err
	}

	return item, nil
}
