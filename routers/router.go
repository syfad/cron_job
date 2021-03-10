package routers

import (
	"github.com/astaxie/beego"
	"cron_job/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/task", &controllers.TaskController{}, "*:Auditlist")
}
