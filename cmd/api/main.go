package main

import (
	"log"
	"net/http"

	handler "github.com/souz-dev/edger-system-go/internal/handler/http"
	"github.com/souz-dev/edger-system-go/internal/repository/memory"
	"github.com/souz-dev/edger-system-go/internal/service"
)

func main() {
	accountRepo := memory.NewAccountRepository()
	transactionRepo := memory.NewTransactionRepository()

	ledgerService := service.NewLedgerService(accountRepo, transactionRepo)

	accountHandler := handler.NewAccountHandler(ledgerService)
	transactionHandler := handler.NewTransactionHandler(ledgerService)

	http.HandleFunc("/transactions", transactionHandler.CreateTransaction)
	http.HandleFunc("/accounts", accountHandler.CreateAccount)
	http.HandleFunc("/accounts/", accountHandler.GetAccountByID)

	log.Println("server running on :5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
