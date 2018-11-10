package bible

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// NewConnectionMySQL create a connection with a rds
func NewConnectionMySQL(config *Config) (*sql.DB, error) {
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false",
		config.AWSUser, config.AWSPassword, config.MySqlbEndpoint, config.AWSInstance,
	)

	dbconn, err := sql.Open("mysql", dnsStr)

	return dbconn, err
}
