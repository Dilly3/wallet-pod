package api

import (
	"net/http"

	"github.com/dilly3/wallet-pod/app/internal/api/handlers"
	"github.com/dilly3/wallet-pod/app/internal/service"
	middleware "github.com/go-chi/chi/middleware"
	chi "github.com/go-chi/chi/v5"
)

func NewRouter(banker service.Banker) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Route("/api", func(r chi.Router) {
		// wallet routes
		r.Post("/wallets/{id}/deposit", handlers.DepositHandler(banker))
		r.Post("/wallets/{id}/withdraw", handlers.WithdrawHandler(banker))
		r.Post("/wallets/transfer", handlers.TransferHandler(banker))
	})
	return r
}
