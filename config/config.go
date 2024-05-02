package config

import (
	"fmt"

	"github.com/backend-magang/cats-social-media/utils/pkg"
	"github.com/spf13/viper"
)

type Config struct {
	AppHost    string `mapstructure:"APP_HOST"`
	AppPort    string `mapstructure:"APP_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBParams   string `mapstructure:"DB_PARAMS"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	BCryptSalt string `mapstructure:"BCRYPT_SALT"`
	SqlTrx     *pkg.SqlWithTransactionService
}

func Load() (conf Config) {
	viper.SetConfigFile("env")
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return
}

func (cfg *Config) GetDSN() (dsn string) {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s&search_path=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBParams,
		cfg.DBSchema,
	)
}
