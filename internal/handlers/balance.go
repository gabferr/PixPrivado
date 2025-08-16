// Em internal/handlers/balance.go

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pix-privado/internal/database"
	"strconv" // <-- Importe o pacote strconv

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BalanceHandler struct {
	DBClient *database.Client
}

func (h *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountIDStr := vars["id"]

	accountID, err := uuid.Parse(accountIDStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	// A função agora retorna o saldo como uma string para evitar overflow
	balanceStr, err := h.DBClient.GetBalance(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error getting balance:", err)
		return
	}

	// <<< CORREÇÃO APLICADA AQUI >>>
	// Convertemos a string do saldo de volta para int64 para a resposta JSON.
	balanceInt, err := strconv.ParseInt(balanceStr, 10, 64)
	if err != nil {
		// Isso só aconteceria se o valor no banco não fosse um número,
		// o que é improvável, mas é bom tratar o erro.
		http.Error(w, "Internal server error: invalid balance format", http.StatusInternalServerError)
		log.Printf("Error converting balance string '%s' to int64: %v", balanceStr, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{
		"balance": balanceInt, // Usamos o valor convertido para int64
	})
}
