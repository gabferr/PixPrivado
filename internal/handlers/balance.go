package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pix-privado/internal/database"

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

	balance, err := h.DBClient.GetBalance(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Error getting balance:", err)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{
		"balance": balance, // Em centavos
	})
}
