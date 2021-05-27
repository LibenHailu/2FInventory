package query

const (
	// insert a new item to an item table
	QueryInsertItem = "INSERT INTO items ( name,price,code,barcode,expire_date,warning_date,quantity,registered_by) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
)
