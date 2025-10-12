package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	DBHost               string
	DBUser               string
	DBPassword           string
	DBName               string
	DBPort               int
	JWTSecret            []byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
)

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "appdb")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("JWT_SECRET", "defaultsecret")
	viper.SetDefault("ACCESS_TOKEN_DURATION", "15m")
	viper.SetDefault("REFRESH_TOKEN_DURATION", "24h")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables and defaults")
	}

	DBHost = viper.GetString("DB_HOST")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
	DBPort = viper.GetInt("DB_PORT")
	JWTSecret = []byte(viper.GetString("JWT_SECRET"))

	var err error
	AccessTokenDuration, err = time.ParseDuration(viper.GetString("ACCESS_TOKEN_DURATION"))
	if err != nil {
		log.Fatalf("Invalid ACCESS_TOKEN_DURATION: %v", err)
	}

	RefreshTokenDuration, err = time.ParseDuration(viper.GetString("REFRESH_TOKEN_DURATION"))
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_DURATION: %v", err)
	}
}
