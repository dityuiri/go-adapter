package graphql

import "github.com/spf13/viper"

type Configuration struct {
	ClientURL string
}

func NewConfig() *Configuration {
	return &Configuration{
		ClientURL: viper.GetString("GRAPHQL_CLIENT_URL"),
	}
}
