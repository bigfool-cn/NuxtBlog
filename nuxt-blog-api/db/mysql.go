package db

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"nuxt-blog-api/configs"
	"strconv"
)
var Eloquent *gorm.DB

func init() {
	dbType := configs.Configs.Db.DbType
	host := configs.Configs.Db.Host
	port := configs.Configs.Db.Port
	database := configs.Configs.Db.Database
	username := configs.Configs.Db.User
	password := configs.Configs.Db.Pwd
	if dbType != "mysql" {
		fmt.Println("db type unknow")
	}
	var err error

	var conn bytes.Buffer
	conn.WriteString(username)
	conn.WriteString(":")
	conn.WriteString(password)
	conn.WriteString("@tcp(")
	conn.WriteString(host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(database)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

	log.Println(conn.String())

	var db Database
	if dbType == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	Eloquent, err = db.Open(dbType, conn.String())

	if err != nil {
		log.Fatalln("mysql connect error %v", err)
	} else {
		log.Println("mysql connect success!")
	}

	Eloquent.LogMode(true)

	if Eloquent.Error != nil {
		log.Fatalln("database error %v", Eloquent.Error)
	}

}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}