package bible

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config map somes variables defaults from application
type Config struct {
	APP  string `envconfig:"APP_NAME" default:"api-bible"`
	Port int    `envconfig:"APP_PORT" default:"8083"`

	DBRegion         string `envconfig:"DB_REGION" default:"us-east-1"`
	DBInstance       string `envconfig:"DB_INSTANCE"`
	DBUser           string `envconfig:"DB_USER"`
	DBPassword       string `envconfig:"DB_PASSWORD"`
	ConnectionString string `envconfig:"DB_CONNECTION_STRING" default:"%s:%s@tcp(%s)/%s?tls=false"`

	DBEndpoint string `envconfig:"DB_ENDPOINT"`
}

//NewConfig config constructor
func NewConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
