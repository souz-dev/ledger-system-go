package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/souz-dev/edger-system-go/internal/domain"
	"github.com/souz-dev/edger-system-go/internal/service"
)

type TransactionHandler struct {
	ledgerService *service.LedgerService
}

func NewTransactionHandler(ledgerService *service.LedgerService) *TransactionHandler {
	return &TransactionHandler{
		ledgerService: ledgerService,
	}
}

type createTransactionRequest struct {
	ID      string               `json:"id"`
	Name    string               `json:"name"`
	Entries []createEntryRequest `json:"entries"`
}

type createEntryRequest struct {
	ID        string           `json:"id"`
	AccountID string           `json:"account_id"`
	Direction domain.Direction `json:"direction"`
	Amount    int64            `json:"amount"`
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req createTransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	entries := make([]domain.Entry, 0, len(req.Entries))

	for _, entryReq := range req.Entries {
		entry, err := domain.NewEntry(
			entryReq.ID,
			entryReq.AccountID,
			entryReq.Direction,
			entryReq.Amount,
			time.Now(),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		entries = append(entries, entry)
	}

	transaction, err := domain.NewTransaction(
		req.ID,
		req.Name,
		entries,
		time.Now(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.ledgerService.CreateTransaction(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
