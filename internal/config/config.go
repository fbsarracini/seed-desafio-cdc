package config

import "os"

type Config struct {
	Addr        string
	DatabaseURL string
}

// .env config
func Load() Config {
	return Config{
		Addr:        getEnv("ADDR", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", "casadocodigo.db"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
