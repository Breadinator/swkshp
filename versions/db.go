package versions

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"

	"github.com/breadinator/swkshp/config"
)

func GetDBPath(game string, create ...bool) (string, error) {
	p, ok := config.Conf.Games[game]
	if !ok {
		return "", fmt.Errorf("game %s not found in %+v", game, config.Conf.Games)
	}

	dbPath := filepath.Join(p, DBName)
	var err error
	if len(create) >= 1 && create[0] {
		err = createDBFromPath(dbPath)
	}

	return dbPath, err
}

func dbExec(game, query string, args ...any) (sql.Result, error) {
	db, err := DBOpen(game)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func dbQuery(game, query string, args ...any) (*sql.Rows, error) {
	db, err := DBOpen(game)
	if err != nil {
		return nil, err
	}
	return db.Query(query, args...)
}

func createDBFromPath(path string) error {
	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		f.Close()
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS mods (id INTEGER UNIQUE, path STRING UNIQUE, sum BLOB, updated DATETIME)`)
	return err
}

func UpdateModEntry(game string, entry Entry) (sql.Result, error) {
	return dbExec(game, `INSERT OR REPLACE INTO mods (id, path, sum, updated) VALUES (?, ?, ?, ?)`, entry.ID, entry.Path, entry.Sum, entry.Updated)
}

func GetModEntry(game string, id int) (*Entry, error) {
	rows, err := dbQuery(game, `SELECT * FROM mods WHERE id=? LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var e Entry
	return &e, scan.Row(&e, rows)
}

func RemoveModEntry(game string, id int) (sql.Result, error) {
	return dbExec(game, "DELETE FROM mods WHERE id=?", id)
}

func GetAllEntries(game string) ([]Entry, error) {
	db, err := DBOpen(game)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT * FROM mods`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	scanned := make([]Entry, 0)
	return scanned, scan.Rows(&scanned, rows)
}
