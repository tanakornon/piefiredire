package main

import (
	"os"
	"piefiredire/internal/adapter/delivery/handler"
	"piefiredire/internal/adapter/repositories/api"
	"piefiredire/internal/core/services"
	"piefiredire/internal/middleware"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type AppConfig struct {
	Port     string
	Endpoint string
	Timeout  int
}

func main() {
	godotenv.Load("./etc/local.env")

	var config AppConfig

	config.Port = os.Getenv("PORT")
	config.Endpoint = os.Getenv("BACON_IPSUM_URL")
	config.Timeout, _ = strconv.Atoi(os.Getenv("TIMEOUT"))

	e := echo.New()

	// beefRepo := inmemory.NewBeefRepository()
	beefRepo := api.NewBeefRepository(config.Endpoint)
	beefService := services.NewBeefService(beefRepo)
	beefHandler := handler.NewBeefHandler(beefService)

	e.Use(
		middleware.NewCustomMiddleware(),
		middleware.NewLoggerMiddleware(),
		middleware.NewTimeoutMiddleware(config.Timeout),
	)

	beefHandler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(config.Port))
}
