package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dilly3/wallet-pod/app/internal/service"
	"github.com/go-chi/chi/v5"
)

type WalletHandler struct {
	Service service.Banker
}

func NewWalletHandler(s service.Banker) *WalletHandler {
	return &WalletHandler{Service: s}
}

type depositRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Reference   *string `json:"reference,omitempty"`
}

type withdrawRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Reference   *string `json:"reference,omitempty"`
}

type transferRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	ToWalletID  int     `json:"to_wallet_id"`
	Reference   *string `json:"reference,omitempty"`
}

// DepositHandler handles POST /wallets/{id}/deposit
func DepositHandler(s service.Banker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, _ := strconv.Atoi(chi.URLParam(r, "id"))
		var req depositRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.Deposit(r.Context(), walletID, req.Amount, req.Description, req.Reference); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "deposit successful"})
	}
}

// WithdrawHandler handles POST /wallets/{id}/withdraw
func WithdrawHandler(s service.Banker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		walletID, _ := strconv.Atoi(chi.URLParam(r, "id"))
		var req withdrawRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.Withdraw(r.Context(), walletID, req.Amount, req.Description, req.Reference); err != nil {
			if err == service.ErrInsufficientFunds {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "withdraw successful"})
	}
}

// TransferHandler handles POST /wallets/transfer
func TransferHandler(s service.Banker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req transferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fromIDStr := r.URL.Query().Get("from_id")
		fromID, _ := strconv.Atoi(fromIDStr)

		if err := s.Transfer(r.Context(), fromID, req.ToWalletID, req.Amount, req.Description, req.Reference); err != nil {
			if err == service.ErrInsufficientFunds {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "transfer successful"})
	}
}
