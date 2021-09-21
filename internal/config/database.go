package config

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var errEndpointDatabaseIsEmpty = errors.New("endpoint database can`t be is empty")

func NewDatabase(getter *viper.Viper) *sqlx.DB {
	endpoint := getter.GetString("database.endpoint")
	if endpoint == "" {
		panic(errEndpointDatabaseIsEmpty)
	}

	db, err := sqlx.Connect("postgres", endpoint)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func (c *config) Database() *sqlx.DB {
	return c.database
}
