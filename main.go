package main

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/gazimaksudur2/BeegoMVC/config"
)

func main() {
	if _, err := config.Load(); err != nil {
		log.Fatal(err)
	}
	beego.Run()
}
