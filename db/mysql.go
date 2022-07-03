package db

import (
	"golang2022/config"
	"golang2022/module"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func MysqlInit() {
	conf := config.Conf
	dsn := conf.Mysql.User + ":" + conf.Mysql.Password + "@tcp(" + conf.Mysql.Host + ":" + conf.Mysql.Port + ")/" + conf.Mysql.Database + "?charset=utf8&parseTime=True&loc=Local"

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	var person module.Person
	result := db.First(&person)
	log.Printf("%+v", result)
	log.Printf("%+v", person)

}
