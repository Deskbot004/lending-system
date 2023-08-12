package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"lending-system/ent"

	"entgo.io/ent/dialect"

	"modernc.org/sqlite"
)

/*
The following struct, func, func are a workaorund to get ent working without gcc in go
*/
type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		err = conn.Close()
		return nil, fmt.Errorf("IDK: %v", err)
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

/*
SetupDB starts a connection to the database

return: client that requested the database; Throws error
*/
func SetupDB() (*ent.Client, error) {
	client, err := ent.Open(dialect.SQLite, "file:db/game_db.db?cache=shared")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, nil
}
