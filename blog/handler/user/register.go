package user

import (
	"blog/middleware"
	"blog/model"
	pkg "blog/pkg"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//@Summary "登录"
//@Description "登录"
//@Tags user
//@Accept application/json
//@Produce application/json
//@Param user body model.User true "账号密码"
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /post/person/register [post]
func Register(c *gin.Context) {
	var user_New = new(model.User)
	var user = new(model.User)
	var user_Old = new(model.User)
	var star = new(model.Star)
	var con = new(model.Control)

	if err := c.ShouldBindJSON(user_New); err != nil {
		log.Println(err, "bind err")
		pkg.Send(c, 500, "should bind json throw an error")
		return
	}

	// 获取最近一条，进行id++
	// user.Latest()
	err := con.SelectFunc(user, "", "latest")

	if err == gorm.ErrRecordNotFound {
		(*user_New).Id = 1
		log.Println("第一个用户产生")
	} else {
		(*user_New).Id = (*user).Id + 1
	}

	// 根据传入的user获取，获取成功，则已存在
	// user_Old.Select((*user_New).Account)
	err = con.SelectFunc(user_Old, (*user_New).Account, "select")
	if err != nil {
		log.Println(err)
		pkg.Send(c, 500, "select throw an error")
		return
	}

	// 对同一个user使用获取信息，他不会进行覆盖
	if (*user_Old).Account == (*user_New).Account {
		pkg.Send(c, 200, "该账号已存在")
		return
	}
	log.Println(user, user_New, user_Old)
	(*user_New).Login_time = time.Now().Format("2006-01-02 15:04:05")

	// user_New.Create()
	err = con.SelectFunc(user_New, "", "create")
	if err != nil {
		log.Println(err, "save error")
		pkg.Send(c, 500, "error in database")
		return
	}

	token, err := middleware.GenerateToken((*user_New).Account)
	if err != nil {
		log.Println(err, "toekn generate error")
		pkg.Send(c, 500, "token failed to generate")
		return
	}

	(*star).Account = (*user_New).Account
	(*star).Time = (*user_New).Login_time

	// star.Create()
	err = con.SelectFunc(star, "", "create")
	if err != nil {
		log.Println(err, "database error")
		pkg.Send(c, 500, "database error")
		return
	}

	c.JSON(200, pkg.Resp{
		Code: 200,
		Msg:  "注册成功",
		Data: token,
	})
}
