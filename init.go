package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subramanian0331/urlShortener/handlers"
	"github.com/subramanian0331/urlShortener/shortenService"
	"github.com/subramanian0331/urlShortener/store"
	"log"
	"os"
)

type App struct {
	fiber *fiber.App
}

// Init bootstrapping the code
func (a *App) Init() {
	a.fiber = fiber.New()
	db, err := store.NewDynamoDB(
		os.Getenv("DB_URL"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_REGION"))
	if err != nil {
		log.Fatal(err.Error())
	}
	svc := shortenService.NewShortenService(db, os.Getenv("APP_URL"), os.Getenv("APP_PORT"))
	h := handlers.NewBaseHandler(svc)
	a.InitRoutes(h)
}

// InitRoutes initializing routes
func (a *App) InitRoutes(h *handlers.BaseHandler) {
	a.fiber.Post("/shorten", h.ShortenUrlHandler)
	a.fiber.Get("/:shortCode", h.RedirectUrlHandler)
}
