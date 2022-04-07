package versions

import (
	"database/sql"
	"errors"
	"path/filepath"

	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3"

	"github.com/breadinator/swkshp/config"
)

func GetDBPath(game string, create ...bool) (string, error) {
	p, ok := config.GetGame(game)
	if !ok {
		return "", errors.New("no game found")
	}

	dbPath := filepath.Join(p, DBName)
	var err error
	if len(create) >= 1 && create[0] {
		err = CreateDBFromPath(dbPath)
	}

	return dbPath, err
}

// Shouldn't keep opening and closing the db, but for the scope of this project shouldn't be problematic.
// TODO find a better way of doing this without bloating
func dbExec(dbPath, query string, args ...any) (sql.Result, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func dbQuery(dbPath, query string, args ...any) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return db.Query(query, args...)
}

func CreateDBFromPath(path string) error {
	_, err := dbExec(path, `CREATE TABLE IF NOT EXISTS mods (id INTEGER UNIQUE, path STRING UNIQUE, sum BLOB, updated DATETIME)`)
	return err
}

func CreateDBFromGame(game string) error {
	p, ok := config.GetGame(game)
	if !ok {
		return errors.New("no game found")
	}

	return CreateDBFromPath(filepath.Join(p, DBName))
}

func UpdateModEntry(game string, entry Entry) (sql.Result, error) {
	p, err := GetDBPath(game, true)
	if err != nil {
		return nil, err
	}

	return dbExec(p, `INSERT OR REPLACE INTO mods (id, path, sum, updated) VALUES (?, ?, ?, ?)`, entry.ID, entry.Path, entry.Sum, entry.Updated)
}

func GetModEntry(game string, id int) (*Entry, error) {
	p, err := GetDBPath(game, true)
	if err != nil {
		return nil, err
	}

	rows, err := dbQuery(p, `SELECT * FROM mods WHERE id=? LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var e Entry
	return &e, scan.Row(&e, rows)
}

func RemoveModEntry(game string, id int) (sql.Result, error) {
	p, err := GetDBPath(game)
	if err != nil {
		return nil, err
	}

	return dbExec(p, "DELETE FROM mods WHERE id=?", id)
}
