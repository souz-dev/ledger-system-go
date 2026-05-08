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
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req createAccountRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
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
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.ledgerService.CreateAccount(account); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, account)
}

func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/accounts/")

	account, err := h.ledgerService.GetAccountByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, account)
}
