package main

import (
	"fmt"
	"golang2022/config"
	"golang2022/db"
)

func main() {
	fmt.Println("Hello, world!")
	config.ViperInit()
	db.MysqlInit()
}

func init() {
	fmt.Println("init")
}
