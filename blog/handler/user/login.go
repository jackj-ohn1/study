package user

import (
	"blog/middleware"
	"blog/model"
	pkg "blog/pkg"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

//@Summary "登录"
//@Description "登录"
//@Tags user
//@Accept application/json
//@Produce application/json
//@Param user body model.User true "账号密码"
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /post/person/login [post]
func Login(c *gin.Context) {
	var user = new(model.User)
	var user_Old = new(model.User)
	var con = new(model.Control)

	if err := c.ShouldBindJSON(user); err != nil {
		log.Println(err, "bind err")
		pkg.Send(c, 500, "should bind json throw an error")
		return
	}

	// user_Old.Select((*user).Account);
	if err := con.SelectFunc(user_Old, (*user).Account, "select"); err != nil {
		log.Println(err, "select error")
		pkg.Send(c, 500, "select error")
		return
	}

	if (*user).Account != (*user_Old).Account {
		log.Println("账号不存在")
		pkg.Send(c, 200, "账号不存在")
		return
	}

	if (*user).Password != (*user_Old).Password {
		pkg.Send(c, 200, "密码错误")
		return
	}

	(*user).Login_time = time.Now().Format("2006-01-02 15:04:05")

	// user.Update((*user).Account)
	err := con.SelectFunc(user, (*user).Account, "update")
	if err != nil {
		pkg.Send(c, 500, "database error")
		log.Println("save 失败", err)
		return
	}

	token, err := middleware.GenerateToken((*user_Old).Account)
	if err != nil {
		log.Println(err, "toekn generate error")
		pkg.Send(c, 500, "token failed to generate")
		return
	}

	c.JSON(200, pkg.Resp{
		Data: token,
		Msg:  "登陆成功",
		Code: 200,
	})

}
