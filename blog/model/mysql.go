package model

import (
	"blog/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = config.Name + ":" + config.PWD + "@tcp" + config.IP + "/blog?" + config.Var
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func DatabaseInit() {
	if err != nil {
		log.Println(err, "database init err")
		return
	}
}

type DatabaseOperation interface {
	Select(account string) error // 获取指定的记录
	Update(account string) error // 保存记录
	Latest() error               // 获取最新的一条记录
	Create() error
}

func (con *Control) SelectFunc(operation DatabaseOperation, account string, flag string) error {
	var err error
	switch flag {
	case "select":
		err = operation.Select(account)
	case "update":
		err = operation.Update(account)
	case "latest":
		err = operation.Latest()
	case "create":
		err = operation.Create()
	}
	return err
}

// user的方法
func (user *User) Latest() error {
	return DB.Model(&User{}).Last(user).Error
}

func (user *User) Update(account string) error {
	return DB.Model(&User{}).Where("account", account).Updates((*user)).Error
}

func (user *User) Create() error {
	return DB.Model(&User{}).Create(user).Error
}

func (user *User) Select(account string) error {
	return DB.Model(&User{}).Where("account=?", account).Find(user).Error
}

// star的方法
func (star *Star) Latest() error {
	return DB.Model(&Star{}).Last(star).Error
}

func (star *Star) Update(account string) error {
	return DB.Model(&Star{}).Where("account=?", account).Updates((*star)).Error
}

func (star *Star) Select(account string) error {
	return DB.Model(&Star{}).Where("account=?", account).Find(star).Error
}

func (star *Star) Create() error {
	return DB.Model(&Star{}).Create(star).Error
}

// 指定更新一项，可以更新零值
func (star *Star) Force(account string) error {
	return DB.Model(&Star{}).Where("account=?", account).Save(star).Error
}

/*
func (user *User) Force(account string) error {
	return DB.Model(&User{}).Where("account=?", account).Save(user).Error
}

func (user *User) Select_Take(account string) error {
	return DB.Model(&User{}).Where("account=?", account).Take(user).Error
}*/
