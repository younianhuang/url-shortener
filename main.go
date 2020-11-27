package main

import (
	_ "url-shortener/routers"

	beego "github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
