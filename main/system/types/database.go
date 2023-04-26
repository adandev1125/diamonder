package types

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	UseDatabase bool
	Driver      string
	Username    string
	Password    string
	Host        string
	Port        int
	Database    string
	ParseTime   bool
}

type Database struct {
	Config DatabaseConfig
	DB     *sql.DB
}

func (d *Database) Connect() {
	var err error
	d.DB, err = sql.Open(
		d.Config.Driver,
		fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?parseTime=%t",
			d.Config.Username, d.Config.Password, d.Config.Host,
			d.Config.Port, d.Config.Database, d.Config.ParseTime,
		),
	)

	if err != nil {
		log.Fatal("Database Connect Error:", err)
	}

	err = d.DB.Ping()

	if err != nil {
		log.Fatal("Database Ping Error:", err)
	}
}
