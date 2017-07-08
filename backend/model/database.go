package model

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/motopig/gatekeeper/backend/common"
)

var (
	err error
	db  *gorm.DB
)

func InitMysql() {

	database := common.Hconfig.String("mysql::database")
	user := common.Hconfig.String("mysql::user")
	password := common.Hconfig.String("mysql::password")
	host := common.Hconfig.String("mysql::host")
	port := common.Hconfig.String("mysql::port")
	param := "?"
	loc := url.QueryEscape("Asia/Shanghai")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%sloc=%s&charset=utf8&parseTime=true", user, password, host, port, database, param, loc)

	db, err = gorm.Open("mysql", conn)

	if err != nil {
		panic(err.Error())
	}

}
