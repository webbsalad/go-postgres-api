package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"aws-0-eu-west-2.pooler.supabase.com"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"postgres.cxywrdpnsfqlylxtkrdf"`
	Password string `env:"DB_PASSWORD" env-default:"AMUHPlFSVqtWqBCc"`
}

func LoadConfig() (ConfigDatabase, error) {
	var cfg ConfigDatabase
	err := cleanenv.ReadEnv(&cfg)
	return cfg, err
}
