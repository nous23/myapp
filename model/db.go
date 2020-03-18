package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var db *sql.DB

type dbConfig struct {
	User     string
	Password string
	Protocol string
	Host     string
	Port     int
	DBName   string
}

func init() {
	err := initDB()
	if err != nil {
		panic(err)
	}
}

func initDBConfig(configPath string) (*dbConfig, error) {
	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("read db config file %s failed: %v\n", configPath, err)
		return nil, err
	}
	dc := dbConfig{}
	err = json.Unmarshal(b, &dc)
	if err != nil {
		log.Printf("parse db config failed: %v\n", err)
		return nil, err
	}
	return &dc, nil
}

func initDB() error {
	currPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("get current path failed: %v\n", err)
		return err
	}
	dbConfigPath := filepath.Join(currPath, "conf", "db", "dbconfig.json")
	dc, err := initDBConfig(dbConfigPath)
	if err != nil {
		return err
	}
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dc.User, dc.Password, dc.Protocol, dc.Host, dc.Port,
		dc.DBName)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("open db failed: %v\n", err)
		return err
	}
	return nil
}
