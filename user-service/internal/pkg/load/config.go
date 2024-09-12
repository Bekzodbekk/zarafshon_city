package load

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Service  ServiceConfig
	Postgres PostgresConfig
}

func LoadConfig(path string) (*Config, error) {

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Service: ServiceConfig{
			Host: viper.GetString("user-service.host"),
			Port: viper.GetString("user-service.port"),
		},
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetString("postgres.port"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
			Database: viper.GetString("postgres.dbname"),
		},
	}
	return &cfg, nil

}