package main

import (
	"context"
	"log"
	"time"

	"github.com/vavelour/warehouse/warehouse/pkg/warehouse_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	warehouseID = 1
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v.", err)
	}
	defer conn.Close()

	c := warehouse_v1.NewWarehouseV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetRemainingItems(ctx, &warehouse_v1.GetRemainingItemsRequest{WarehouseId: warehouseID})
	if err != nil {
		log.Printf("Failed to get items by id: %v", err)
	}

	log.Printf("Items: %v", r.GetItemList())
}
