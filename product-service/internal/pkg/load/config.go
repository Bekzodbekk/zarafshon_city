package load

import "github.com/spf13/viper"

type Mongosh struct {
	Host       string
	Port       int
	Database   string
	Collection string
}

type Config struct {
	Mongosh Mongosh

	ProductServiceHost string
	ProductServicePort int
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	conf := Config{
		Mongosh: Mongosh{
			Host:       viper.GetString("mongosh.host"),
			Port:       viper.GetInt("mongosh.port"),
			Database:   viper.GetString("mongosh.database"),
			Collection: viper.GetString("mongosh.collection"),
		},

		ProductServiceHost: viper.GetString("service.host"),
		ProductServicePort: viper.GetInt("service.port"),
	}

	return &conf, nil
}
