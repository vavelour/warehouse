package service

import (
	"context"

	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/models"
	"github.com/vavelour/warehouse/warehouse/internal/service/warehouse/mapper"
)

type WarehouseRepository interface {
	GetItems(ctx context.Context, id int64) ([]models.ItemModel, error)
	ReservedItems(ctx context.Context, reservations []models.ReservationModel) (bool, error)
	ReleaseReserve(ctx context.Context, reservations []models.ReservationModel) (bool, error)
}

type WarehouseService struct {
	repo WarehouseRepository
}

func NewWarehouseService(repo WarehouseRepository) *WarehouseService {
	return &WarehouseService{repo: repo}
}

func (s *WarehouseService) GetItems(ctx context.Context, id int64) ([]entities.Item, error) {
	items, err := s.repo.GetItems(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ItemModelToEntities(items), nil
}

func (s *WarehouseService) ReservedItems(ctx context.Context, reservations []entities.Reservation) (bool, error) {
	success, err := s.repo.ReservedItems(ctx, mapper.ReservationEntitiesToModel(reservations))
	if err != nil {
		return success, err
	}

	return success, nil
}

func (s *WarehouseService) ReleaseReserve(ctx context.Context, reservations []entities.Reservation) (bool, error) {
	success, err := s.repo.ReleaseReserve(ctx, mapper.ReservationEntitiesToModel(reservations))
	if err != nil {
		return success, err
	}

	return success, nil
}
