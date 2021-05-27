package grpcserver

import (
	"context"
	"log"

	"github.com/LibenHailu/2FInventory/cmd/item/itempb"
	"github.com/LibenHailu/2FInventory/internal/constant/model"
	"github.com/LibenHailu/2FInventory/internal/module/item"
)

type GrpcServer struct {
	ItemUsecase item.Usecase
	itempb.UnimplementedItemServiceServer
}

func (g *GrpcServer) CreateItem(ctx context.Context, req *itempb.CreateItemRequest) (*itempb.CreateItemResponse, error) {

	newItem := req.GetItem()

	data := &model.Item{
		Name:         newItem.GetName(),
		Price:        newItem.GetPrice(),
		Code:         newItem.GetCode(),
		Barcode:      newItem.GetBarcode(),
		ExpireDate:   newItem.GetExpireDate(),
		WarningDate:  newItem.GetWarningDate(),
		Quantity:     newItem.GetQuantity(),
		RegisteredBy: newItem.GetRegisteredBy(),
	}

	item, err := g.ItemUsecase.InsertItem(ctx, data)

	if err != nil {
		log.Fatalf("some thing went wrong: %V", err.Error())
	}

	return &itempb.CreateItemResponse{
		Item: &itempb.Item{
			Name:         item.Name,
			Price:        item.Price,
			Code:         item.Code,
			Barcode:      item.Barcode,
			ExpireDate:   item.ExpireDate,
			WarningDate:  item.WarningDate,
			Quantity:     item.Quantity,
			RegisteredBy: item.RegisteredBy,
		},
	}, nil
}
