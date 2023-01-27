package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost            string `mapstructure:"db_host"`
	DBPort            int    `mapstructure:"db_port"`
	DBUser            string `mapstructure:"db_user"`
	DBName            string `mapstructure:"db_name"`
	DBNameForTest     string `mapstructure:"db_name_for_test"`
	DBPassword        string `mapstructure:"db_password"`
	MidtransClientKey string `mapstructure:"midtrans_client_key"`
	MidtransServerKey string `mapstructure:"midtrans_server_key"`
}

func LoadConfigFromEnv() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return &config
}

func (config *Config) CreateDBURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}
