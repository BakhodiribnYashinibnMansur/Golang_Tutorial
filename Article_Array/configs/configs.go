package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	App         string
	Environment string // dev,test, prod

	ServiceHost      string
	HTTPPort         string
	Version          string
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	DefaultOffset string
	DefaultLimit  string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{
		App:          cast.ToString(getOrReturnDefaultValue("PROJECT_NAME", "go_boilerplate")),
		Environment:  cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "dev")),
		Version:      cast.ToString(getOrReturnDefaultValue("VERSION", "1.0")),
		ServiceHost:  cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhsot")),
		HTTPPort:     cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080")),
		PostgresHost: cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost")),
		PostgresPort: cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432)),

		PostgresUser:     cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres")),
		PostgresPassword: cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "your_db_password")),
		PostgresDatabase: cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "your_db_database")),
		DefaultOffset:    cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0")),
		DefaultLimit:     cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10")),
	}

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
