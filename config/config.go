package config

import (
	"makeup/analytics/logger"
	"os"

	"github.com/Valgard/godotenv"
)

func Init() {
	dotenv := godotenv.New()

	envFiles := []string{
		".env",
		".env.local",
		".env.test",
	}

	if os.Getenv("SERVER_PORT") == "" {
		os.Setenv("SERVER_PORT", "9000")
	}

	for _, envFile := range envFiles {
		if _, err := os.Stat(envFile); err == nil {
			if err := dotenv.Overload(envFile); err != nil {
				logger.StdErr("Failed to read " + envFile + " file: " + err.Error())
				panic(err)
			}
		}
	}
}
