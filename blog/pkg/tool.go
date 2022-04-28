package pkg

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Send(c *gin.Context, code int, msg string) {
	c.JSON(code, Resp{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func Get() ([]string, error) {
	path := "./MdFile/"
	var content = make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return content, err
	}

	for _, file := range files {
		name := strings.Replace(file.Name(), ".md", "", -1)
		content = append(content, name)
	}
	return content, nil
}
