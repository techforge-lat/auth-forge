package localconfig

import (
	"auth-forge/internal/shared/domain"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/techforge-lat/errortrace/v2"
)

var envVars []string

func clearPreviousEnv() {
	for _, key := range envVars {
		os.Unsetenv(key)
	}
	envVars = []string{}
}

func loadEnvFile(path string) error {
	env := os.Getenv("ENV")
	if env != "local" {
		return nil
	}

	clearPreviousEnv()

	envMap, err := godotenv.Read(path)
	if err != nil {
		return err
	}

	for key, value := range envMap {
		os.Setenv(key, value)
		envVars = append(envVars, key)
	}

	return nil
}

func getEnvOrFail(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errortrace.OnError(fmt.Errorf("missing %s environment variable", key))
	}

	return value, nil
}

func parseUint(value string, key string) (uint64, error) {
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, errortrace.OnError(fmt.Errorf("invalid value for %s: %w", key, err))
	}

	return parsed, nil
}

func NewLocalConfig(path string) (domain.Configuration, error) {
	_ = loadEnvFile(path)

	portHTTP, err := parseUint(os.Getenv("PORT_HTTP"), "PORT_HTTP")
	if err != nil {
		return domain.Configuration{}, errortrace.OnError(err)
	}

	dbPort, err := parseUint(os.Getenv("DB_PORT"), "DB_PORT")
	if err != nil {
		return domain.Configuration{}, errortrace.OnError(err)
	}

	dbEngine, err := getEnvOrFail("DB_ENGINE")
	if err != nil {
		return domain.Configuration{}, errortrace.OnError(err)
	}

	config := domain.Configuration{
		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
		AllowedMethods: strings.Split(os.Getenv("ALLOWED_METHODS"), ","),
		Env:            os.Getenv("ENV"),
		PortHTTP:       uint(portHTTP),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		Timezone:       os.Getenv("TIMEZONE"),
		Database: domain.Database{
			Driver:   dbEngine,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_SERVER"),
			Port:     uint(dbPort),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		OTEL: domain.OTEL{
			CollectorEndpoint: os.Getenv("OTEL_COLLECTOR_ENDPOINT"),
		},
	}

	return config, nil
}
