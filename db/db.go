package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

//config for server
type serverConfig struct {
	User         string `json:"mysql-user"`
	Password     string `json:"mysql-password"`
	DatabaseName string `json:"db-name"`
	HostName     string `json:"host"`
	Port         int    `json:"port"`
}

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	//setup configuration
	config := loadConfiguration()
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.User, config.Password, config.HostName, config.Port, config.DatabaseName))
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		err = db.Ping()
	}
	return db, err
}

func loadConfiguration() serverConfig {
	file := "./config.json"
	var config serverConfig
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
