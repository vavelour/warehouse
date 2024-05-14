package mapper

import (
	"github.com/vavelour/warehouse/warehouse/internal/domain/entities"
	"github.com/vavelour/warehouse/warehouse/pkg/warehouse_v1"
)

func EntitiesItemToProtobufItem(items []entities.Item) []*warehouse_v1.Item {
	sizeCap := len(items)
	protobufItems := make([]*warehouse_v1.Item, 0, sizeCap)

	for _, item := range items {
		protobufItem := &warehouse_v1.Item{
			Name:       item.Name,
			Size:       int64(item.Size),
			UniqueCode: int64(item.UniqueCode),
			Quantity:   int64(item.Quantity),
		}
		protobufItems = append(protobufItems, protobufItem)
	}

	return protobufItems
}
