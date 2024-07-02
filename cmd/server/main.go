package main

import (
	"pokemon-api-server/internal/db"
	"pokemon-api-server/internal/pokemon/model"
	"pokemon-api-server/internal/pokemon/router"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db.ConnectDB()

	model.GetPokemonTypeMap()
	model.GetPokemonGenerationMap()

	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PATCH},
	}))

	router.AttachPokemonRoutes(e)
	router.AttachGenerationRoutes(e)
	router.AttachTypeRoutes(e)

	router.AttachTestRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8000"))
}
