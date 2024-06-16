package config

import "fmt"

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "1234"
	DBName     = "postgres"
	DBSSLMode  = "disable"
)

func GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", DBUser, DBPassword, DBHost, DBPort, DBName, DBSSLMode)
}
