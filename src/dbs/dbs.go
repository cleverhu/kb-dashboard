package dbs

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
	"os"
	"time"
)

var (
	Orm *gorm.DB
)

func init() {
	err := fmt.Errorf("")
	Orm, err = gorm.Open("mysql", "devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mysqlDB := Orm.DB()
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxIdleConns(5)
	Orm.LogMode(true)

}
