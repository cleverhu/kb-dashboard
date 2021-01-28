package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	Orm *gorm.DB
)

func init() {
	fmt.Println("init db")
	Orm = gormDB()
	//fmt.Println(Orm)
}

func gormDB() *gorm.DB {
	//var newLogger = logger.New(
	//	// io.writer同样使用colorable
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 开启彩色打印
	//	},
	//)
	orm, err := gorm.Open(mysql.Open("devuser:123~!@@tcp(39.105.28.235:3320)/tech?charset=utf8mb4&parseTime=true&loc=Local"),
		&gorm.Config{Logger: nil})//
	if err != nil {
		log.Fatal(err)
	}
	mysqlDB, _ := orm.DB()
	mysqlDB.SetConnMaxLifetime(30 * time.Second)
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)


	return orm
}
