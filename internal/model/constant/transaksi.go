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
	NoRek        string  `json:"norek"`
	Gram         float64 `json:"gram"`
	HargaTopUp   int64   `json:"harga"`
	HargaBuyBack int64   `json:"-"`
}

type BuyBackRequest struct {
	NoRek        string  `json:"norek"`
	Gram         float64 `json:"gram"`
	HargaTopUp   int64   `json:"_"`
	HargaBuyBack int64   `json:"harga"`
}
