package controllers

import (
	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

func (c *TaskController) Auditlist() {
	c.Data["pageTitle"] = "任务管理"
	c.Layout = "public/layout.html"
	c.TplName = "task/auditlist.html"
}