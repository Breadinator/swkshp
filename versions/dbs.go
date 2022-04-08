package versions

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var databases = make(map[string]*sql.DB)

func DBOpen(game string) (*sql.DB, error) {
	db, ok := databases[game]
	if ok {
		return db, nil
	}

	path, err := GetDBPath(game, true)
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("sqlite3", path)
	if err == nil {
		databases[game] = db
	}
	return db, err
}

func DBCloseAll() (errs []error) {
	for game := range databases {
		if err := databases[game].Close(); err != nil {
			errs = append(errs, err)
		} else {
			delete(databases, game)
		}
	}
	return
}

func DBLen() int {
	return len(databases)
}
