package models

type ReservationModel struct {
	UniqueCode  int `db:"item_warehouse_id"`
	WarehouseId int `db:"warehouse_id"`
	Quantity    int `db:"quantity"`
}
