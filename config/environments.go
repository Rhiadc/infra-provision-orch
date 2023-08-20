package config

import (
	"fmt"

	"github.com/pingcap/log"
	"github.com/spf13/viper"
)

type Environments struct {
	APIPort        string
	ENV            string
	Git            Git
	GitConfigChart GitConfigChart
}

type GitConfigChart struct {
	RepoURL           string
	RepoInternalPath  string
	RepoReferenceName string
	Username          string
	Password          string
	OwnerName         string
	OnwerEmail        string
}

type Git struct {
	Token              string
	InfratemplateToken string
	Owner              string
	GolangTemplateRepo string
	InfraTemplateRepo  string
	APIVersion         string
	GlobalToken        string
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("ENV", "dev")
	viper.SetDefault("API_PORT", "8080")
	viper.SetDefault("REPO_OWNER", "Rhiadc")
	viper.SetDefault("GOLANG_TEMPLATE_REPO", "https://www.github.com/rhiadc/")
	viper.SetDefault("TOKEN", "some-token")
	viper.SetDefault("API_VERSION", "v1")
	viper.SetDefault("INFRA_TEMPLATE_REPO", "https://www.github.com/rhiadc/")
	viper.SetDefault("INFRA_TEMPLATE_TOKEN", "some-token")
	viper.SetDefault("GLOBAL_TOKEN", "some-token")

	//GitConfigChart
	viper.SetDefault("REPO_URL", "https://www.github.com/rhiadc/")
	viper.SetDefault("REPO_INTERNAL_PATH", "repo")
	viper.SetDefault("REPO_REFERENCE_NAME", "MAIN")
	viper.SetDefault("USERNAME", "Rhiadc")
	viper.SetDefault("PASSWORD", "*")
	viper.SetDefault("OWNER_EMAIL", "rhiad.ciccoli@gmail.com")

	if err := viper.ReadInConfig(); err != nil {
		log.Info(fmt.Sprintf("unable to find or read config file: %s", err))
	}

	return &Environments{
		APIPort: viper.GetString("API_PORT"),
		ENV:     viper.GetString("ENV"),
		Git: Git{
			Token:              viper.GetString("TOKEN"),
			GolangTemplateRepo: viper.GetString("GOLANG_TEMPLATE_REPO"),
			InfratemplateToken: viper.GetString("INFRA_TEMPLATE_TOKEN"),
			GlobalToken:        viper.GetString("GLOBAL_TOKEN"),
			InfraTemplateRepo:  viper.GetString("INFRA_TEMPLATE_REPO"),
			Owner:              viper.GetString("REPO_OWNER"),
			APIVersion:         viper.GetString("API_VERSION"),
		},
		GitConfigChart: GitConfigChart{
			RepoURL:           viper.GetString("REPO_URL"),
			RepoInternalPath:  viper.GetString("REPO_INTERNAL_PATH"),
			RepoReferenceName: viper.GetString("REPO_REFERENCE_NAME"),
			Username:          viper.GetString("USERNAME"),
			Password:          viper.GetString("PASSWORD"),
			OnwerEmail:        viper.GetString("OWNER_EMAIL"),
			OwnerName:         viper.GetString("REPO_OWNER"),
		},
	}
}
