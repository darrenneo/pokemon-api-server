package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"pokemon-api-server/internal/db"
	"pokemon-api-server/internal/middleware"
	gormModel "pokemon-api-server/internal/model"
	"pokemon-api-server/internal/pokemon/model"
	"pokemon-api-server/internal/query"

	"github.com/labstack/echo/v4"
)

func GetPokemonById(c echo.Context) error {
	pokeId := c.Param("id")
	pokemonID64, error1 := strconv.ParseInt(pokeId, 10, 64)
	if error1 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error1.Error()})
	}
	pokemonDb := query.Use(db.DB).Pokemon
	pokemonArr, error2 := pokemonDb.Where(pokemonDb.ID.Eq(pokemonID64)).Find()
	if error2 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error2.Error()})
	}

	editGenTypeField(pokemonArr)

	return c.JSON(http.StatusOK, pokemonArr)
}

func GetPokemonByDexId(c echo.Context) error {
	pokeId := c.Param("id")
	pokemonID64, err := strconv.ParseInt(pokeId, 10, 64)
	if err != nil {
		panic(err)
	}

	pokemonDb := query.Use(db.DB).Pokemon
	pokemonArr, err := pokemonDb.Where(pokemonDb.PokedexNumber.Eq(pokemonID64)).Find()
	if err != nil {
		panic(err)
	}

	editGenTypeField(pokemonArr)

	return c.JSON(http.StatusOK, pokemonArr)
}

func GetPokemonByGeneration(c echo.Context) error {
	pokeGeneration := c.Param("generation")

	pokeGeneration64, error1 := strconv.ParseInt(pokeGeneration, 10, 64)
	if error1 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error1.Error()})
	}

	genString := model.GenMap[int(pokeGeneration64)]

	pokemonDb := query.Use(db.DB).Pokemon

	allPokemon, error2 := pokemonDb.Where(pokemonDb.Generation.Eq(pokeGeneration64)).Find()
	if error2 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error2.Error()})
	}

	editGenTypeField(allPokemon)

	return c.JSON(http.StatusOK, model.PokemonResponse{Pokemons: allPokemon, Generations: genString})
}

func GetPokemonByGenerationArray(c echo.Context) error {
	pokeGeneration := c.Param("generation")

	pokeGeneration64, error1 := strconv.ParseInt(pokeGeneration, 10, 64)
	if error1 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error1.Error()})
	}

	pokemonDb := query.Use(db.DB).Pokemon

	allPokemon, error2 := pokemonDb.Where(pokemonDb.Generation.Eq(pokeGeneration64)).Find()
	if error2 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error2.Error()})
	}

	editGenTypeField(allPokemon)

	return c.JSON(http.StatusOK, allPokemon)
}

func GetPokemonByType(c echo.Context) error {
	typeName := c.Param("type")
	var typeId int64
	for k, v := range model.TypeMap {
		if strings.Contains(strings.ToLower(v), strings.ToLower(typeName)) {
			typeId = int64(k)
		}
	}

	pokemonDb := query.Use(db.DB).Pokemon
	pokemonArr, err := pokemonDb.Where(pokemonDb.Type1.Eq(typeId)).Or(pokemonDb.Type2.Eq(typeId)).Find()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	editGenTypeField(pokemonArr)

	return c.JSON(http.StatusOK, pokemonArr)
}

func SearchPokemon(c echo.Context) error {
	fmt.Print("Running this\n")
	pokeQuery := c.QueryParam("q")

	pokemonDb := query.Use(db.DB).Pokemon
	pokemonArr, error2 := pokemonDb.Where(pokemonDb.Name.Lower().Like("%" + pokeQuery + "%")).Find()
	if error2 != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: error2.Error()})
	}

	editGenTypeField(pokemonArr)

	return c.JSON(http.StatusOK, pokemonArr)
}

func RefetchGen(c echo.Context) error {
	err := model.GetPokemonGenerationMap()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	fmt.Print("Refreshing Gen\n")

	return c.JSON(http.StatusOK, "Gen Updated")
}

func RefetchType(c echo.Context) error {
	err := model.GetPokemonTypeMap()
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}
	fmt.Print("Refreshing Type\n")
	return c.JSON(http.StatusOK, "Type Updated")
}

func editGenTypeField(pokeArr []*gormModel.Pokemon) {
	for _, pokemon := range pokeArr {
		pokemon.Type1String = model.TypeMap[int(pokemon.Type1)]
		if pokemon.Type2 != 0 {
			pokemon.Type2String = model.TypeMap[int(pokemon.Type2)]
		}
	}
}

func RandomMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

/////////////////////////////////////////// With Pagination ///////////////////////////////////////////

func GetAllPokemons(c echo.Context) error {
	pagination := c.Get("pagination").(*middleware.PaginationStruc)
	// x := c.Get("per_page").(int64)
	allPokemon, err := db.GetAllPokemons(c.Request().Context(), pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	editGenTypeField(allPokemon)

	return c.JSON(http.StatusOK, allPokemon)
}

func GetAllTypes(c echo.Context) error {
	pagination := c.Get("pagination").(*middleware.PaginationStruc)

	allTypes, err := db.GetAllTypes(c.Request().Context(), pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, allTypes)
}

func GetAllGeneration(c echo.Context) error {
	pagination := c.Get("pagination").(*middleware.PaginationStruc)

	allGen, err := db.GetAllGeneration(c.Request().Context(), pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, allGen)
}

func FilterTry(c echo.Context) error {
	queryParam := c.Get("filter")

	// fmt.Println(*queryParam.(map[string]middleware.FilterStruc)["att"].IntValue.Gte)
	// fmt.Println(*queryParam.(map[string]middleware.FilterStruc)["att"].IntValue.Lt)

	return c.JSON(http.StatusOK, queryParam)
}

func PaginationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		per_page, err := strconv.ParseInt(c.QueryParam("per_page"), 10, 0)
		if err != nil {
			per_page = 10
		}
		page, err := strconv.ParseInt(c.QueryParam("page"), 10, 0)
		if err != nil {
			page = 1
		}

		c.Set("per_page", per_page)
		c.Set("page", page)

		return next(c)
	}
}
