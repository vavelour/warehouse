package models

type ItemModel struct {
	Name       string `db:"name"`
	Size       int    `db:"size"`
	UniqueCode int    `db:"unique_code"`
	Quantity   int    `db:"quantity"`
}
