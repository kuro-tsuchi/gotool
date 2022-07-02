package db

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var MsqlDB *sqlx.DB
var GoquDialect *goqu.Database

func MysqlInit() *sqlx.DB {
	var err error
	MsqlDB, err = sqlx.Connect("mysql", "root:root@(localhost:3306)/test")
	if err != nil {
		log.Fatalln(err)
	}
	return MsqlDB
}

func  DialectInit() *goqu.Database {
	   GoquDialect =   goqu.New("mysql", MsqlDB)
	   return  GoquDialect
}
