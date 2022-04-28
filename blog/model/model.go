package model

type User struct {
	Account    string `json:"account" binding:"required"`
	Id         int    `json:"id"`
	Password   string `json:"password" gorm:"password" binding:"required"`
	Login_time string `gorm:"login_time"`
}

type Star struct {
	Name    string `json:"name" gorm:"name"`
	Time    string
	Account string
}

type Control struct{}
