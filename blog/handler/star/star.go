package star

import (
	"blog/model"
	pkg "blog/pkg"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//@Summary "收藏"
//@Description "收藏"
//@Tags star
//@Accept application/json
//@Produce application/json
//@Param star query string true "star"
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /star/addition [get]
func AddStar(c *gin.Context) {
	var star = new(model.Star)
	var con = new(model.Control)
	var star_Slice []string
	var star_name = c.Query("star")

	content, err := pkg.Get()

	if err != nil {
		log.Println(err)
		pkg.Send(c, 200, "发生了错误")
		return
	}

	if FindOne(star_name, content) == -1 {
		pkg.Send(c, 200, "不存在该博文")
		return
	}

	time := time.Now().Format("2006-01-02 15:04:05")

	account, exists := c.MustGet("id").(string)

	if !exists {
		log.Println("获取id失败")
		pkg.Send(c, 401, "获取id失败")
		return
	}

	// star.Select(account)
	if err := con.SelectFunc(star, account, "select"); err != nil {
		pkg.Send(c, 500, "database error")
		return
	}

	star_all := (*star).Name
	if star_all == "" {
		star_Slice = []string{}
	} else {
		star_Slice = strings.Split(star_all, ",") //
	}

	fmt.Println(star_Slice, len(star_Slice), star_name)

	if FindOne(star_name, star_Slice) == -1 {
		star_Slice = append(star_Slice, star_name)
	} else {
		pkg.Send(c, 200, "已收藏过该博文了")
		return
	}
	fmt.Println(star_Slice)

	(*star).Name = strings.Join(star_Slice, ",")

	(*star).Time = time

	// star.Updata((*star).Account)
	err = con.SelectFunc(star, (*star).Account, "update")

	if err != nil {
		log.Println("database error")
		pkg.Send(c, 500, "database error")
		return
	}

	pkg.Send(c, 200, "add star successfully")
}

//@Summary "收藏"
//@Description "收藏"
//@Tags star
//@Accept application/json
//@Produce application/json
//@Param star query string true "star"
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /star/delete [get]
func DeleteStar(c *gin.Context) {
	var star = new(model.Star)
	var con = new(model.Control)
	var star_name = c.Query("star")
	time := time.Now().Format("2006-01-02 15:04:05")

	account, exists := c.MustGet("id").(string)

	if !exists {
		log.Println("获取id失败")
		pkg.Send(c, 401, "获取id失败")
		return
	}

	// star.Select(account)
	if err := con.SelectFunc(star, account, "select"); err != nil {
		pkg.Send(c, 500, "database error")
		return
	}

	star_all := (*star).Name

	star_Slice := strings.Split(star_all, ",")

	if len(star_Slice) == 1 {
		star_all = ""
	} else {
		key := FindOne(star_name, star_Slice)
		if key == -1 {
			log.Println("不存在该博文")
			pkg.Send(c, 200, "不存在该博文")
			return
		}

		star_Slice = append(star_Slice[:key], star_Slice[(key+1):]...)
		star_all = strings.Join(star_Slice, ",")
	}

	(*star).Time = time
	(*star).Name = star_all

	if star_all == "" {
		err := star.Force(account)
		if err != nil {
			log.Println("database error")
			pkg.Send(c, 500, "database error")
			return
		}
	} else {
		// star.Update((*star).Account)
		err := con.SelectFunc(star, (*star).Account, "update")
		if err != nil {
			log.Println("database error")
			pkg.Send(c, 500, "database error")
			return
		}
	}

	pkg.Send(c, 200, "delete star successfully")
}

func FindOne(target string, all []string) int {
	for k, v := range all {
		if v == target {
			return k
		}
	}
	return -1
}
