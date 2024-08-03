package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Database struct {
	DBHost    string `mapstructure:"host"`
	DBPort    string `mapstructure:"port"`
	DBName    string `mapstructure:"name"`
	DBUser    string `mapstructure:"user"`
	DBPass    string `mapstructure:"password"`
	SSLModels string `mapstructure:"ssl_mode"`
	TimeZone  string `mapstructure:"timezone"`
}

type Email struct {
	SMTPHost   string `mapstructure:"host"`
	SMTPPass   string `mapstructure:"password"`
	SMTPPort   int    `mapstructure:"port"`
	SMTPEmail  string `mapstructure:"email"`
	SMTPSender string `mapstructure:"sender"`
}

type OAuth struct {
	GoogleClientID     string `mapstructure:"google_client_id"`
	GoogleClientSecret string `mapstructure:"google_client_secret"`
}

type Config struct {
	PORT         string        `mapstructure:"port"`
	ClientOrigin string        `mapstructure:"client_origin"`
	TokenSecret  string        `mapstructure:"secret_key"`
	TokenMaxAge  int           `mapstructure:"max_age"`
	TokenExpired time.Duration `mapstructure:"expire_time"`
	OTPExpired   time.Duration `mapstructure:"otp_expire_time"`
	Database     Database      `mapstructure:"database"`
	Email        Email         `mapstructure:"email"`
	OAuth        OAuth         `mapstructure:"oauth"`
}

func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return config, nil
}
