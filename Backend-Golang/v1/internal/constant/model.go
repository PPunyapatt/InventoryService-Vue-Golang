package constant

type Inventories struct {
	InventoryCode string `json:"inventory_code" db:"inventory_code"`
	CreatedAt     string `json:"created_at" db:"created_at"`
}
