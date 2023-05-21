package types

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	UseDatabase bool
	Driver      string
	EnvHost     string
	EnvPort     string
	Username    string
	Password    string
	Database    string
	ParseTime   bool
}

func (r *DBConfig) toString() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=%t",
		r.Username, r.Password, os.Getenv(r.EnvHost), os.Getenv(r.EnvPort), r.Database, r.ParseTime)
}

type Database struct {
	Config DBConfig
	DB     *sql.DB
}

func (d *Database) Connect() {
	var err error
	dbConfig := d.Config.toString()
	log.Printf("Connecting database: %s\n", dbConfig)
	d.DB, err = sql.Open(d.Config.Driver, dbConfig)

	if err != nil {
		log.Fatal("Database Connect Error:", err)
	}

	err = d.DB.Ping()

	if err != nil {
		log.Fatal("Database Ping Error:", err)
	}

	log.Printf("Database Connected.\n\n")
}
