package route

import (
	"net/http"

	"dift_backend_driver/driver-payments-service/internal/handler"
)

func Register(mux *http.ServeMux, h *handler.PaymentsHandler) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/api/v1/driver/wallet", h.GetWallet)
	mux.HandleFunc("/api/v1/driver/wallet/ledger", h.GetWalletLedger)
	mux.HandleFunc("/api/v1/driver/payments/history", h.GetPaymentHistory)
}
