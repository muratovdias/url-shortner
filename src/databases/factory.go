package databases

import (
	"errors"
	"fmt"
	"github.com/muratovdias/url-shortner/src/config"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	"github.com/muratovdias/url-shortner/src/databases/drivers/sqlite"
)

func New(conf config.DataStore) (ds drivers.DataStore, err error) {
	switch conf.DbName {
	case "sqlite3":
		return sqlite.New(conf), nil
	}

	return nil, errors.New(fmt.Sprintf("can't connect to database: %s unimplemented", conf.DbName))
}
