package config

import (
	"fmt"

	"github.com/pingcap/log"
	"github.com/spf13/viper"
	"google.golang.org/appengine/log"
)

type Environments struct {
	APIPort string
	ENV     string
	Git     Git
}

type Git struct {
	Token              string
	Owner              string
	GolangTemplateRepo string
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("API_PORT", "8080")
	viper.SetDefault("REPO_OWNER", "Rhiadc")
	viper.SetDefault("GOLANG_TEMPLATE_REPO", "https://www.github.com/rhiadc/so")
	viper.SetDefault("TOKEN", "some-token")
	if err := viper.ReadInConfig(); err != nil {
		log.Info(fmt.Sprintf("unable to find or read config file: %w", err))
	}

	return &Environments{
		APIPort: viper.GetString("API_PORT"),
		ENV:     viper.GetString("ENV"),
		Git: Git{
			Token:              viper.GetString("TOKEN"),
			GolangTemplateRepo: viper.GetString("GOLANG_TEMPLATE_REPO"),
			Owner:              viper.GetString("REPO_OWNER"),
		},
	}
}
