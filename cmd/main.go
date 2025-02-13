package main

import (
	"avito-shop/config"
	"avito-shop/migrations"
	"avito-shop/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitDB()

	migrations.Migrate(config.DB)

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	fmt.Println("🚀 Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
