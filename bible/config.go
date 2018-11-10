package bible

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config map somes variables defaults from application
type Config struct {
	APP          string `envconfig:"APP_NAME" default:"bible"`
	Port         int    `envconfig:"APP_PORT" default:"8083"`
	FileCreds    string `envconfig:"" default:"/home/user/.aws/credentials"`
	ProfileCreds string `envconfig:"" default:"personal"`

	AWSRegion   string `envconfig:"BIBLE_AWS_REGION" default:"us-east-2"`
	AWSInstance string `envconfig:"BIBLE_AWS_INSTANCE"`
	AWSUser     string `envconfig:"BIBLE_AWS_DB_USER"`
	AWSPassword string `envconfig:"BIBLE_AWS_DB_PASSWORD"`

	MySqlbEndpoint string `envconfig:"BIBLE_AWS_DB_ENDPOINT"`
}

//NewConfig config constructor
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
