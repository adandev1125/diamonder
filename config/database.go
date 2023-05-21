package config

import (
	"main/system/types"
)

// TODO allow multiple databases

/**
 * A database config for the app.
 *
 * @param	UseDatabase		True if this application use a database, false if not.
 * @param	Driver			Database driver name. If you use other databases like MongoDB,
 *							you must import the driver in system/database/database.go.
 * @param	Username		Username used to connect to the database.
 * @param	PasswordEnv		Environment variable name for password used to connect to the database.
 * @param	Host			The database host.
 * @param	Port			The port of database server.const.
 * @param	Database		The database name used for this app.
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
