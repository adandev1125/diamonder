package config

import (
	"main/main/system/types"
)

var DBConfig types.DatabaseConfig = types.DatabaseConfig{
	Driver:    "mysql",
	Username:  "root",
	Password:  "",
	Host:      "127.0.0.1",
	Port:      3306,
	Database:  "test",
	ParseTime: true,
}
