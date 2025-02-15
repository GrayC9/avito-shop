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
	mux.HandleFunc("/api/buy/{item}", handlers.BuyMerchHandler)
	mux.HandleFunc("/api/sendCoin", handlers.SendCoinsHandler)

	mux.HandleFunc("/api/info", handlers.InfoHandler)
}
