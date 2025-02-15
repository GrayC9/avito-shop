package routes

import (
	"avito-shop/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "pong"}`))
	})

	mux.HandleFunc("/api/auth", handlers.TokenHandler)
	//mux.HandleFunc("/revoke", handlers.RevokeHandler)

	mux.HandleFunc("/api/buy/", handlers.BuyMerchHandler)
	mux.HandleFunc("/api/sendCoin", handlers.SendCoinsHandler)
}
