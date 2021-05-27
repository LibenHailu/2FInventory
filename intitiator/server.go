package intitiator

import (
	"log"
	"net"

	grpcserver "github.com/LibenHailu/2FInventory/cmd/item/grpc_server"
	"github.com/LibenHailu/2FInventory/cmd/item/itempb"
	"github.com/LibenHailu/2FInventory/internal/module/item"
	"github.com/LibenHailu/2FInventory/internal/storage/persistance"
	"github.com/LibenHailu/2FInventory/platform/postgres"
	"google.golang.org/grpc"
)

func Inventory(testInit bool) {
	databaseItem := persistance.ItemInit(postgres.Conn)
	itemUsecase := item.Initialize(databaseItem)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v ", err)
	}

	s := grpc.NewServer()
	itempb.RegisterItemServiceServer(s, &grpcserver.GrpcServer{
		ItemUsecase: itemUsecase,
	})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
