package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	base_db *sql.DB
}

func Open(db_path string) (*Database, error) {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return nil, err
	}
	var count int
	if err = db.QueryRow(`
	SELECT count(*) 
	FROM sqlite_master 
	WHERE type="table" AND name="userinfo"
	`).Scan(&count); err != nil {
		return nil, err
	}
	if count == 0 {
		if _, err = db.Exec(`
		CREATE TABLE userinfo (
			uid INTEGER PRIMARY KEY NOT NULL,
			username text NULL,
			password text NULL,
			created DATE NULL
		);
		`); err != nil {
			return nil, err
		}
	}
	return &Database{db}, nil
}

func (db *Database) Close() {
	db.base_db.Close()
}

func (db *Database) GetUser(uid int64) (string, string, error) {
	var username, password string
	err := db.base_db.QueryRow(`
		SELECT username, password
		FROM userinfo
		WHERE uid = ?
	`, uid).Scan(&username, &password)
	return username, password, err
}

func (db *Database) AddUser(uid int64, username, password string) error {
	_, err := db.base_db.Exec(`
	INSERT INTO userinfo (
		uid,
		username,
		password,
		created
	) 
	VALUES (?, ?, ?, datetime('now', 'localtime'))
	`, uid, username, password)
	return err
}

func (db *Database) SetUser(uid int64, username, password string) error {
	_, _, err := db.GetUser(uid)
	if err == sql.ErrNoRows {
		return db.AddUser(uid, username, password)
	}
	_, err = db.base_db.Exec(`
	UPDATE userinfo
	SET username = ?, password = ?, created = datetime('now', 'localtime')
	WHERE uid = ?
	`, username, password, uid)
	return err
}

func (db *Database) DeleteUser(uid int64) error {
	_, err := db.base_db.Exec(`
	DELETE FROM userinfo
	WHERE uid = ?
	`, uid)
	return err
}
