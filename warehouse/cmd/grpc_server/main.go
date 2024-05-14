package main

import (
	"fmt"
	"log"
	"net"

	server "github.com/vavelour/warehouse/warehouse/internal/api/warehouse"
	"github.com/vavelour/warehouse/warehouse/internal/config"
	repository "github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/postgres"
	repos "github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/repos"
	service "github.com/vavelour/warehouse/warehouse/internal/service/warehouse"
	"github.com/vavelour/warehouse/warehouse/pkg/warehouse_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = "50051"

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	db, err := repository.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	warehouseRepo := repos.NewWarehouseRepos(db)
	warehouseService := service.NewWarehouseService(warehouseRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	warehouse_v1.RegisterWarehouseV1Server(s, server.NewWarehouseServer(warehouseService))

	log.Printf("Server listening at port: %s", grpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
