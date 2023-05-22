package system

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Name      string `yaml:"name"`       // Name of the database
	Driver    string `yaml:"driver"`     // Database driver to use
	Host      string `yaml:"host"`       // Hostname or IP address of the database server
	Port      string `yaml:"port"`       // Port number to connect to the database server
	Database  string `yaml:"database"`   // Name of the database to connect to
	Username  string `yaml:"username"`   // Username to authenticate with
	Password  string `yaml:"password"`   // Password to authenticate with
	ParseTime bool   `yaml:"parse_time"` // Whether or not to parse time fields in the database
}

func (r *DbConfig) toString() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=%t",
		r.Username, r.Password, r.Host, r.Port, r.Database, r.ParseTime)
}

type Config struct {
	Port      int        `yaml:"port"`      // Port number for the server
	Name      string     `yaml:"name"`      // Name of the server
	DbConfigs []DbConfig `yaml:"databases"` // List of database configurations
}

// Initialize Config
func (r *Config) Init() {
	// Open Config File
	file, error := os.Open("./config/config.yml")
	if error != nil {
		panic(error)
	}
	// Close file after function exits
	defer file.Close()
	// Check if file is not empty
	if file == nil {
		panic("unable to open the config.yml")
	}
	// Decode YAML data
	decoder := yaml.NewDecoder(file)
	error = decoder.Decode(r)
	if error != nil {
		panic(error)
	}
	// Connect to databases if there are configurations available
	if len(r.DbConfigs) > 0 {
		r.ConnectDatabases()
	}
}

// ConnectDatabases connects to databases defined in Config struct
func (r *Config) ConnectDatabases() {
	// Initialize map to store databases
	databases := make(map[string]*sql.DB)
	// Loop through database configurations and initialize connections
	for _, dbConfig := range r.DbConfigs {
		// Convert database configuration to string for connection
		configStr := dbConfig.toString()
		// Log connection attempt
		log.Printf("Connecting database: %s\n", configStr)
		// Open database connection
		db, err := sql.Open(dbConfig.Driver, configStr)
		if err != nil {
			// Panic if unable to open connection
			panic(err)
		}
		// Ping database to verify connection
		err = db.Ping()
		if err != nil {
			// Panic if unable to ping database
			panic(err)
		}
		// Store database connection in map
		databases[dbConfig.Name] = db
		// Log successful connection
		log.Printf("Database Connected.\n\n")
	}
}

var config *Config
var databases map[string]*sql.DB
var configOnce sync.Once

func GetConfig() *Config {
	configOnce.Do(func() {
		config = &Config{}
	})

	return config
}

func GetDatabase(name string) *sql.DB {
	db, ok := databases[name]
	if ok {
		return db
	} else {
		return nil
	}
}

func CloseDatabases() {
	for name, db := range databases {
		db.Close()
		delete(databases, name)
	}
}
