package config

import "golang_graphs/backend/internal/database"

type Config struct {
	Host string `config:"APP_HOST" yaml:"host"`
	Port string `config:"APP_PORT" yaml:"port"`

	Postgres     database.Config `config:"postgres"`
	TestPostgres database.Config `config:"test_postgres"`
}
