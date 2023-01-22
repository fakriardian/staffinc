package constant

type TransactionType string

const (
	TransactionTypeBuyBack = "buyback"
	TransactionTypeTopUp   = "topup"
)

// like schema
type CheckMutasiRequest struct {
	NoRek     string `json:"norek"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}

type TopUpRequest struct {
	NoRek      string `json:"norek"`
	Gram       int64  `json:"gram"`
	HargaTopUp int64  `json:"harga"`
}

type BuyBackRequest struct {
	NoRek        string `json:"norek"`
	Gram         int64  `json:"gram"`
	HargaBuyBack int64  `json:"harga"`
}
