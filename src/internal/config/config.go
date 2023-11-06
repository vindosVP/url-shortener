package config

import "github.com/caarlos0/env/v8"

type Config struct {
	Env            string `env:"ENV,notEmpty" envDefault:"local"`
	Domain         string `env:"DOMAIN,notEmpty"`
	DomainProtocol string `env:"DOMAIN_PROTOCOL,notEmpty"`
	StorageType    string `env:"STORAGE_TYPE,notEmpty"`
	Server         Server
	DB             DB
}

type Server struct {
	Type     string `env:"SERVER_TYPE"`
	APIPort  string `env:"API_PORT" envDefault:"8080"`
	GRPCPort string `env:"GRPC_PORT" envDefault:"8081"`
}

type DB struct {
	Host     string `env:"DB_HOST" `
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER"`
	Pwd      string `env:"DB_PWD"`
	Name     string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	TimeZone string `env:"DB_TIMEZONE" envDefault:"Europe/Moscow"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
