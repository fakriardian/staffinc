package model

type Harga struct {
	AdminID      string `json:"admin_id"`
	HargaTopUp   int64  `json:"harga_topup"`
	HargaBuyBack int64  `json:"harga_buyback"`
}
