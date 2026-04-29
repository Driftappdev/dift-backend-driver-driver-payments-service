package handler

import (
	"net/http"

	response "github.com/driftappdev/libpackage/contracts/response"

	"dift_backend_driver/driver-payments-service/internal/service"
)

type PaymentsHandler struct{ svc *service.PaymentsService }

func NewPaymentsHandler(svc *service.PaymentsService) *PaymentsHandler {
	return &PaymentsHandler{svc: svc}
}

func (h *PaymentsHandler) GetWallet(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{
			Error: &response.AppError{Code: "bad_request", Message: "driver_id required"},
		})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetWallet(driverID)})
}

func (h *PaymentsHandler) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{
			Error: &response.AppError{Code: "bad_request", Message: "driver_id required"},
		})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetPaymentHistory(driverID)})
}

func (h *PaymentsHandler) GetWalletLedger(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		writeJSON(w, http.StatusBadRequest, response.Envelope[any]{
			Error: &response.AppError{Code: "bad_request", Message: "driver_id required"},
		})
		return
	}
	writeJSON(w, http.StatusOK, response.Envelope[any]{Data: h.svc.GetWalletLedger(driverID)})
}
