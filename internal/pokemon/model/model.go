package model

import (
	"pokemon-api-server/internal/db"
	"pokemon-api-server/internal/model"
	dbModel "pokemon-api-server/internal/model"
	"pokemon-api-server/internal/query"
)

var TypeMap map[int]string

var GenMap map[int]model.StringSlice

type PokemonResponse struct {
	Pokemons    []*dbModel.Pokemon `json:"pokemons"`
	Generations []string           `json:"generation,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetPokemonTypeMap() error {
	pokemonTypeMap := make(map[int]string)
	typeDb := query.Use(db.DB).PokemonType
	pokemonType, err := typeDb.Find()
	if err != nil {
		return err
	}

	for _, pokeType := range pokemonType {
		pokemonTypeMap[int(pokeType.ID)] = pokeType.Name
	}

	TypeMap = pokemonTypeMap

	return nil
}

func GetPokemonGenerationMap() error {
	pokemonGenerationMap := make(map[int]model.StringSlice)
	generationDb := query.Use(db.DB).Generation
	pokemonGeneration, err := generationDb.Find()
	if err != nil {
		return err
	}

	for _, genName := range pokemonGeneration {
		pokemonGenerationMap[int(genName.Generation)] = genName.VideoGameTitles
	}

	GenMap = pokemonGenerationMap

	return nil
}
