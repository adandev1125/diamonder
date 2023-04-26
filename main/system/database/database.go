package database

import (
	"database/sql"
	"main/main/config"
	"main/main/system/types"
	"sync"
)

var database *types.Database
var once sync.Once

func GetDatabase() *sql.DB {
	once.Do(func() {

		database = &types.Database{
			Config: config.DBConfig,
		}

		if config.DBConfig.UseDatabase {
			database.Connect()
		}

	})
	return database.DB
}
