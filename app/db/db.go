package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"sync"
)

var (
	connection gorm.DB
	once       sync.Once
)

func New() gorm.DB {
	once.Do(Init)
	return connection
}

func Init() {
	database, err := gorm.Open("mysql", getConnectionString())

	if err != nil {
		revel.ERROR.Printf("DATABASE CONNECTION ERROR: %v", err)
	}

	connection = database
}

func getConnectionString() string {
	var (
		user     string
		password string
		table    string
	)

	user, ok := revel.Config.String("mysql.user")
	if !ok {
		panic("Config variable mysql.user not defined")
	}
	password, ok = revel.Config.String("mysql.password")
	if !ok {
		panic("Config variable mysql.password not defined")
	}
	table, ok = revel.Config.String("mysql.db")
	if !ok {
		panic("Config variable mysql.db not defined")
	}
	return fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", user, password, table)
}
