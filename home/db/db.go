package database

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct{ *sql.DB }

func Open(dbname string) (db DB, err error) {
	db.DB, err = sql.Open("sqlite3", dbname)
	if err != nil {
		return
	}

	err = db.Ping()
	return
}

func (db DB) NewMatch(uid, key string) (err error) {
	stmt, err := db.Prepare("INSERT INTO Clients VALUES(?, ?)")
	if err != nil {
		return
	}

	_, err = stmt.Exec(uid, key)
	return
}

type Client struct {
	UID string
	Key string
}

func (db DB) GetClient(uid string) (c Client, err error) {
	row := db.QueryRow("SELECT * FROM Matches WHERE UID = ?", uid)
	err = row.Scan(c.UID, c.Key)
	return
}
