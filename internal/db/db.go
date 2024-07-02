package db

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"pokemon-api-server/internal/middleware"
	"pokemon-api-server/internal/model"
	"pokemon-api-server/internal/query"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDbUrl() string {
	const secretFile = "secrets.json"
	var data map[string]string
	unformattedUrl, err := os.ReadFile(secretFile)
	if err != nil {
		panic(err)
	}

	err1 := json.Unmarshal([]byte(unformattedUrl), &data)
	if err1 != nil {
		panic(err1)
	}

	value, ok := data["database_url"]
	if !ok {
		panic("Failed to get database url")
	}

	return value
}

func ConnectDB() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
		},
	)
	DatabaseURL := GetDbUrl()
	db, err := gorm.Open(postgres.Open(DatabaseURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetAllPokemons(ctx context.Context, pagination *middleware.PaginationStruc) ([]*model.Pokemon, error) {
	pokemonDb := query.Use(DB).Pokemon.WithContext(ctx)

	if pagination != nil {
		pokemonDb = pokemonDb.Offset(int(pagination.Page-1) * int(pagination.PerPage)).Limit(int(pagination.PerPage))
	}

	return pokemonDb.Find()
}

func GetAllTypes(ctx context.Context, pagination *middleware.PaginationStruc) ([]*model.PokemonType, error) {
	typeDb := query.Use(DB).PokemonType.WithContext(ctx)

	if pagination != nil {
		typeDb = typeDb.Offset(int(pagination.Page-1) * int(pagination.PerPage)).Limit(int(pagination.PerPage))
	}

	return typeDb.Find()
}

func GetAllGeneration(ctx context.Context, pagination *middleware.PaginationStruc) ([]*model.Generation, error) {
	generationDb := query.Use(DB).Generation.WithContext(ctx)

	if pagination != nil {
		generationDb = generationDb.Offset(int(pagination.Page-1) * int(pagination.PerPage)).Limit(int(pagination.PerPage))
	}

	return generationDb.Find()
}

// func generateFilterConditions(filter interface{}) []gen.Condition {

// }
