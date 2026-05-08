package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/souz-dev/edger-system-go/internal/domain"
	"github.com/souz-dev/edger-system-go/internal/service"
)

type AccountHandler struct {
	ledgerService *service.LedgerService
}

func NewAccountHandler(ledgerService *service.LedgerService) *AccountHandler {
	return &AccountHandler{
		ledgerService: ledgerService,
	}
}

type createAccountRequest struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Direction domain.Direction `json:"direction"`
	Balance   int64            `json:"balance"`
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req createAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	account, err := domain.NewAccount(
		req.ID,
		req.Name,
		req.Direction,
		req.Balance,
		time.Now(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.ledgerService.CreateAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/accounts/")

	account, err := h.ledgerService.GetAccountByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}
