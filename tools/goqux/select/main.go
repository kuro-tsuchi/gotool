package main

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"goqusqlx/db"
	"goqusqlx/module"
	"log"
)

func main() {
	log.Println("main")
	// use dialect.From to get a dataset to build your SQL
	//ds := db.GoquDialect.From("person").Where(goqu.Ex{"first_name": "Bin"})
	//sql, args, err := ds.ToSQL()
	//if err != nil {
	//	fmt.Println("An error occurred while generating the SQL", err.Error())
	//}
	//fmt.Println(sql, args)

	//var person module.Person
	var persons []module.Person

	// get data
	//err = db.MsqlDB.Get(&person, sql)
	//fmt.Printf("%#v\n", person)

	// select data
	//toSQL, _, err := db.GoquDialect.From("person").Where(goqu.Ex{"first_name": []string{"John", "Bin", "Jane"}}).Limit(5).ToSQL()
	dialect := goqu.Dialect("mysql")
	toSQL, args, _ := dialect.From("person").Where(goqu.C("first_name").ILike("%I%")).ToSQL()
	fmt.Printf("%#v\n", toSQL)
	fmt.Printf("%#v\n", args)
	db.MsqlDB.Select(&persons, toSQL)
	for _, val := range persons {
		fmt.Printf("val:%#v\n", val)

	}

}
func init() {
	db.DialectInit()
	db.MysqlInit()
}
