package config

import "github.com/spf13/viper"

type WarehouseConfig struct {
	DB DBConfig
}

func InitConfig() (WarehouseConfig, error) {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return WarehouseConfig{}, err
	}

	cfg := WarehouseConfig{
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			User:     viper.GetString("db.user"),
			DBName:   viper.GetString("db.db_name"),
			Password: viper.GetString("db.password"),
			SSLMode:  viper.GetString("db.ssl_mode"),
		},
	}

	return cfg, nil
}
