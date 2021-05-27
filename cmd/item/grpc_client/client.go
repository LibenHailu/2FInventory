package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LibenHailu/2FInventory/cmd/item/itempb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not found connect %v ", err)
	}

	defer cc.Close()

	c := itempb.NewItemServiceClient(cc)

	callInsertItem(c)

}

func callInsertItem(c itempb.ItemServiceClient) {

	req := &itempb.CreateItemRequest{
		Item: &itempb.Item{
			Name:         "cake",
			Price:        50,
			Code:         "C10",
			Barcode:      "A1056asfh5300",
			ExpireDate:   "10-15-2023",
			WarningDate:  "10-15-2023",
			Quantity:     130,
			RegisteredBy: 23,
		},
	}
	res, err := c.CreateItem(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response form Great: %v", res.Item)
}
