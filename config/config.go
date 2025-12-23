package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Services ServicesConfig `mapstructure:"services"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type ServicesConfig struct {
	UserService         string `mapstructure:"user_service"`
	ProductService      string `mapstructure:"product_service"`
	OrderService        string `mapstructure:"order_service"`
	PaymentService      string `mapstructure:"payment_service"`
	InventoryService    string `mapstructure:"inventory_service"`
	NotificationService string `mapstructure:"notification_service"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

func LoadConfig() *Config {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, using defaults: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &config
}
