package config

import (
	"main/system/types"
)

// TODO allow multiple databases

/*
*
A database config for the app.
@param	UseDatabase		True if this application use a database, false if not.
@param	Driver			Database driver name.
@param	EnvHost			Environment variable name for the database host.
@param	EnvPort			Environment variable name for the port of database server.
@param	Username		Username used to connect to the database.
@param	Password		Password used to connect to the database.
@param	Database		The database name used for this app.
*/
var DBConfig types.DBConfig = types.DBConfig{
	UseDatabase: true,
	Driver:      "mysql",
	EnvHost:     "MYSQL_HOST",
	EnvPort:     "MYSQL_PORT",
	Username:    "test_user",
	Password:    "test_password",
	Database:    "test",
	ParseTime:   true,
}
