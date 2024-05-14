package mapper

import (
	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/pkg/warehouse_v1"
)

func ReserveItemsRequestToEntities(req *warehouse_v1.ReserveItemsRequest) []entities.Reservation {
	sizeCap := len(req.Products)
	reservations := make([]entities.Reservation, 0, sizeCap)

	for _, product := range req.GetProducts() {
		reservation := entities.Reservation{
			UniqueCode:  int(product.GetUniqueCode()),
			WarehouseId: int(product.GetWarehouseId()),
			Quantity:    int(product.GetQuantity()),
		}
		reservations = append(reservations, reservation)
	}

	return reservations
}

func ReleaseReserveItemsRequestToEntities(req *warehouse_v1.ReleaseReserveItemsRequest) []entities.Reservation {
	sizeCap := len(req.Products)
	reservations := make([]entities.Reservation, 0, sizeCap)

	for _, product := range req.GetProducts() {
		reservation := entities.Reservation{
			UniqueCode:  int(product.GetUniqueCode()),
			WarehouseId: int(product.GetWarehouseId()),
			Quantity:    int(product.GetQuantity()),
		}
		reservations = append(reservations, reservation)
	}

	return reservations
}
