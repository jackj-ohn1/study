package getinfor

import (
	pkg "blog/pkg"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type File struct {
	Name    string
	Content string
}

//@Summary "获取博文"
//@Description "获取博文"
//@Tags blog
//@Accept application/json
//@Produce application/json
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /get/file/data [get]
func GetBlog(c *gin.Context) {
	path := "./MdFile/"
	var contents = make([]File, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		pkg.Send(c, 500, "error")
		return
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		var content File
		name := file.Name()
		words, err := os.ReadFile(path + name)
		if err != nil {
			log.Fatal(err)
		}
		name = strings.Replace(name, ".md", "", -1)
		content.Name = name
		content.Content = string(words)
		contents = append(contents, content)
	}

	if err != nil {
		log.Println(err)
		pkg.Send(c, 500, "error")
		return
	}

	c.JSON(200, pkg.Resp{
		Code: 200,
		Msg:  "successfully",
		Data: contents,
	})
}

func Res(c *gin.Context) {
	c.Redirect(301, "localhost:5500/get/data")
}

//@Summary "获取博文名称"
//@Description "获取博文名称"
//@Tags blog
//@Accept application/json
//@Produce application/json
//@Success 200 {object} pkg.Resp "successfully"
//@Failure 500 {object} pkg.Resp "error happened in server"
//@Router /get/file/name [get]
func GetName(c *gin.Context) {
	content, err := pkg.Get()
	if err != nil {
		pkg.Send(c, 500, "error")
		return
	}

	c.JSON(200, pkg.Resp{
		Code: 200,
		Msg:  "successfully",
		Data: content,
	})
}
