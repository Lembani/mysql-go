package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DefaultCharset = "utf8"
)

type MySQL struct {
	Addr     string
	Username string
	Password string
	DBName   string
	Charset  string
	db       *sql.DB
}

func NewMySQL(addr, username, password, dbName string) *MySQL {
	return &MySQL{
		Addr:     addr,
		Username: username,
		Password: password,
		DBName:   dbName,
		Charset:  DefaultCharset,
	}
}

func (m *MySQL) initDB() (db *sql.DB, err error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=%s`,
		m.Username, m.Password, m.Addr, m.DBName, m.Charset)
	db, err = sql.Open(`mysql`, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
