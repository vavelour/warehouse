package mapper

import (
	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/internal/repository/warehouse/models"
)

func ItemModelToEntities(model []models.ItemModel) []entities.Item {
	sizeCap := len(model)
	items := make([]entities.Item, 0, sizeCap)

	for _, val := range model {
		items = append(items, entities.Item{
			Name:       val.Name,
			Size:       val.Size,
			UniqueCode: val.UniqueCode,
			Quantity:   val.Quantity,
		})
	}

	return items
}
