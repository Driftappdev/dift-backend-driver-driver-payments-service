package service

import (
	"time"

	"dift_backend_driver/driver-payments-service/internal/dto"
)

type PaymentsService struct{}

func NewPaymentsService() *PaymentsService { return &PaymentsService{} }

func (s *PaymentsService) GetWallet(driverID string) dto.WalletResponse {
	ledger := s.GetWalletLedger(driverID)
	return dto.WalletResponse{
		DriverID:       driverID,
		Available:      18450.75,
		PendingPayout:  3200.00,
		CommissionDue:  3180.40,
		LastSettlement: time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
		Ledger:         ledger.Items,
	}
}

func (s *PaymentsService) GetWalletLedger(driverID string) dto.WalletLedgerResponse {
	return dto.WalletLedgerResponse{
		DriverID: driverID,
		Items: []dto.WalletLedgerItem{
			{ID: "wallet-001", Title: "Trip payout", Amount: 425.00, Type: "credit", OccurredAt: time.Now().Add(-3 * time.Hour).Format(time.RFC3339), Description: "3 completed trips"},
			{ID: "wallet-002", Title: "Commission accrual", Amount: 86.00, Type: "debit", OccurredAt: time.Now().Add(-6 * time.Hour).Format(time.RFC3339), Description: "platform commission"},
			{ID: "wallet-003", Title: "Weekly incentive", Amount: 500.00, Type: "credit", OccurredAt: time.Now().Add(-28 * time.Hour).Format(time.RFC3339), Description: "peak hour campaign"},
			{ID: "wallet-004", Title: "Driver reward redemption", Amount: 450.00, Type: "credit", OccurredAt: time.Now().Add(-30 * time.Hour).Format(time.RFC3339), Description: "Peak-hour completions reward redeemed"},
		},
	}
}

func (s *PaymentsService) GetPaymentHistory(driverID string) dto.PaymentHistoryResponse {
	return dto.PaymentHistoryResponse{
		DriverID: driverID,
		Items: []dto.PaymentHistoryItem{
			{ID: "pay-001", Title: "Weekly payout", Type: "payout", Amount: 3200, Status: "settled", OccurredAt: time.Now().Add(-24 * time.Hour).Format(time.RFC3339), Description: "Transferred to Kasikorn Bank"},
			{ID: "pay-002", Title: "Commission debit", Type: "commission", Amount: 860, Status: "posted", OccurredAt: time.Now().Add(-36 * time.Hour).Format(time.RFC3339), Description: "Platform service fee"},
			{ID: "pay-003", Title: "Campaign reward", Type: "reward", Amount: 500, Status: "settled", OccurredAt: time.Now().Add(-72 * time.Hour).Format(time.RFC3339), Description: "Weekend boost incentive"},
		},
	}
}
