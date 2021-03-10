package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id              int
	Group_id        int
	Server_id       int
	Task_name       string
	Description     string
	Cron_spec       string
	Concurrent      int
	Command         string
	Timeout         int
	Execute_times   int
	Prev_time       int
	Is_notify       int
	Notify_type     int
	Notify_user_ids string
	Status          int
	Create_time     int
	Create_id       int
	Update_time     int
	Update_id       int
}

func (task *Task) TableName() string {
	//从配置文件读取，表前缀
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "task"
}

//插入
func (task *Task) Insert() error {
	if _, err := orm.NewOrm().Insert(task); err != nil {
		return err
	}
	return nil
}

//删除
func (task *Task) Delete() error {
	if _, err := orm.NewOrm().Delete(task); err != nil {
		return err
	}
	return nil
}

//更新
func (task *Task) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(task, fields...); err != nil {
		return err
	}
	return nil
}

//读取
func (task *Task) Read(fields ...string) error {
	if err := orm.NewOrm().Read(task, fields...); err != nil {
		return err
	}
	return nil
}
