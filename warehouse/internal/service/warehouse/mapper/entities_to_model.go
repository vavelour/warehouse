package mapper

import (
	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/models"
)

func ReservationEntitiesToModel(reservations []entities.Reservation) []models.ReservationModel {
	sizeCap := len(reservations)
	model := make([]models.ReservationModel, 0, sizeCap)

	for _, val := range reservations {
		model = append(model, models.ReservationModel{
			UniqueCode:  val.UniqueCode,
			WarehouseId: val.WarehouseId,
			Quantity:    val.Quantity,
		})
	}

	return model
}
