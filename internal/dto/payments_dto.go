package dto

type WalletLedgerItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	OccurredAt  string  `json:"occurred_at"`
	Description string  `json:"description"`
}

type WalletResponse struct {
	DriverID       string             `json:"driver_id"`
	Available      float64            `json:"available"`
	PendingPayout  float64            `json:"pending_payout"`
	CommissionDue  float64            `json:"commission_due"`
	LastSettlement string             `json:"last_settlement"`
	Ledger         []WalletLedgerItem `json:"ledger"`
}

type PaymentHistoryItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	OccurredAt  string  `json:"occurred_at"`
	Description string  `json:"description"`
}

type PaymentHistoryResponse struct {
	DriverID string               `json:"driver_id"`
	Items    []PaymentHistoryItem `json:"items"`
}

type WalletLedgerResponse struct {
	DriverID string             `json:"driver_id"`
	Items    []WalletLedgerItem `json:"items"`
}
