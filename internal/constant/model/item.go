package model

type Item struct {
	ID           int64
	Name         string
	Price        int32
	Code         string
	Barcode      string
	ExpireDate   string
	WarningDate  string
	Quantity     int32
	RegisteredBy int64
}
