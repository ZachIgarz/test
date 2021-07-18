package config

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     int    `env:"DATABASE_PORT,default=5432"`
		User     string `env:"DATABASE_USER"`
		Password string `env:"DATABASE_PASSWORD"`
		DbName   string `env:"DATABASE_DB_NAME"`
	}
	Server struct {
		HTTPPort int32 `env:"HTTPPORT,default=3000"`
	}
	
}

var c config

const (
    Test string = "test"
)

//ReadConfig read config
func ReadConfig() error {
	ctx := context.Background()
	err := envconfig.Process(ctx, &c)
	return err
}

/*
//SetConfigs set configs
func SetConfigs(Host string, DbName string, User string, Password string, Migrate bool) error {
	ReadConfig()
	c.Database.Host = Host
	return nil
}*/

// HTTPListener the listener string for the http service
func HTTPListener() string {
	return fmt.Sprintf(":%d", c.Server.HTTPPort)
}

// EnableMigrations get enable migrations
func EnableMigrations() bool {
	return c.Migrate
}
/*
// PgConn the connection string to the pg database
func PgConn() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DbName)
}
*/

/*
// PgConnMigration returns the config string for migration
func PgConnMigration() *string {
	if c.Migrate {
		pgconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			c.Database.User,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.DbName)
		return &pgconn
	}

	return nil
}
*/

