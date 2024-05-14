package server

import (
	"context"

	"github.com/pkg/errors"

	"github.com/vavelour/warehouse/warehouse/internal/api/warehouse/mapper"
	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/pkg/warehouse_v1"
)

type WarehouseService interface {
	GetItems(ctx context.Context, id int64) ([]entities.Item, error)
	ReservedItems(ctx context.Context, reservations []entities.Reservation) (bool, error)
	ReleaseReserve(ctx context.Context, reservations []entities.Reservation) (bool, error)
}

type WarehouseServer struct {
	warehouse_v1.UnimplementedWarehouseV1Server
	service WarehouseService
}

func NewWarehouseServer(s WarehouseService) *WarehouseServer {
	return &WarehouseServer{service: s}
}

func (s *WarehouseServer) GetRemainingItems(ctx context.Context, req *warehouse_v1.GetRemainingItemsRequest) (*warehouse_v1.GetRemainingItemsResponse, error) {
	items, err := s.service.GetItems(ctx, req.GetWarehouseId())
	if err != nil {
		return nil, err
	}

	return &warehouse_v1.GetRemainingItemsResponse{ItemList: mapper.EntitiesItemToProtobufItem(items)}, nil
}

func (s *WarehouseServer) ReserveItems(ctx context.Context, req *warehouse_v1.ReserveItemsRequest) (*warehouse_v1.ReserveItemsResponse, error) {
	reservedResult, err := s.service.ReservedItems(ctx, mapper.ReserveItemsRequestToEntities(req))
	if err != nil {
		return &warehouse_v1.ReserveItemsResponse{Success: reservedResult}, errors.Wrap(err, "Reserve items is failed")
	}

	return &warehouse_v1.ReserveItemsResponse{Success: reservedResult}, nil
}

func (s *WarehouseServer) ReleaseReserveItems(ctx context.Context, req *warehouse_v1.ReleaseReserveItemsRequest) (*warehouse_v1.ReleaseReserveItemsResponse, error) {
	reservedResult, err := s.service.ReleaseReserve(ctx, mapper.ReleaseReserveItemsRequestToEntities(req))
	if err != nil {
		return &warehouse_v1.ReleaseReserveItemsResponse{Success: reservedResult}, errors.Wrap(err, "Release reserve items is failed")
	}

	return &warehouse_v1.ReleaseReserveItemsResponse{Success: reservedResult}, nil
}
