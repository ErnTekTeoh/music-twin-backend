package data

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"music-twin-backend/common"
	"music-twin-backend/config"
	"time"
)

var db *gorm.DB

func init() {

	username := config.GetDBConnection().Username
	password := config.GetDBConnection().Password
	dbHost := config.GetDBConnection().DBHost
	port := config.GetDBConnection().Port
	dbName := config.GetDBConnection().DBName
	loc, _ := time.LoadLocation("Asia/Kuala_Lumpur")

	c := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", dbHost, port),
		DBName:               dbName,
		Loc:                  loc,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	common.LogWithDebug(context.Background(), c.FormatDSN())
	var err error
	db, err = gorm.Open(gormMysql.Open(c.FormatDSN()), &gorm.Config{})
	if err != nil {
		common.LogWithError(context.Background(), err.Error())
		panic(err)
	}
	return
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	fmt.Println("Closing DB connection......")
	conn, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conn.Close()
	fmt.Println(err)
	return
}
