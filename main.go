package main

import (
	_ "cron_job/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

