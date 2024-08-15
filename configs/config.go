package configs

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "${DATABASE_URL}")
	viper.SetDefault("database.port", "${DATABASE_PORT}")
	viper.SetDefault("database.user", "${DB_USER}")
	viper.SetDefault("database.pass", "${DB_PASS}")
}

func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// Substitua as variáveis de ambiente na configuração
	replaceEnvVariables()

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	log.Printf("Database config: %+v\n", cfg.DB) // Adicione este log

	return nil
}

func replaceEnvVariables() {
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envKey := value[2 : len(value)-1]
			envValue := os.Getenv(envKey)
			if envValue != "" {
				viper.Set(key, envValue)
			}
		}
	}
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
