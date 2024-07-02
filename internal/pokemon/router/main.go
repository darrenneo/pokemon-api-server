package router

import (
	"pokemon-api-server/internal/middleware"
	"pokemon-api-server/internal/pokemon/handler"

	"github.com/labstack/echo/v4"
)

func AttachPokemonRoutes(e *echo.Echo) {
	pokemonGroup := e.Group("/pokemon")
	pokemonGroup.Use(middleware.PaginationMiddleWare())
	pokemonGroup.GET("/dex/:id", handler.GetPokemonByDexId)
	pokemonGroup.GET("/:id", handler.GetPokemonById)
	pokemonGroup.GET("", handler.GetAllPokemons)
	pokemonGroup.GET("/gen/:generation", handler.GetPokemonByGenerationArray)
	pokemonGroup.GET("/search", handler.SearchPokemon)
	pokemonGroup.PATCH("/refetchGen", handler.RefetchGen)
	pokemonGroup.PATCH("/refetchType", handler.RefetchType)
}

func AttachGenerationRoutes(e *echo.Echo) {
	genGroup := e.Group("gen")
	genGroup.Use(middleware.PaginationMiddleWare())
	genGroup.GET("", handler.GetAllGeneration)
	genGroup.GET("/:generation", handler.GetPokemonByGeneration)
}

func AttachTypeRoutes(e *echo.Echo) {
	typeGroup := e.Group("/type")
	typeGroup.Use(middleware.PaginationMiddleWare())
	typeGroup.GET("", handler.GetAllTypes)
	typeGroup.GET("/:type", handler.GetPokemonByType)
}

func AttachTestRoutes(e *echo.Echo) {
	e.Use(middleware.FilterMiddleWare())
	e.GET("/test", handler.FilterTry)
}
