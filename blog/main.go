package main

import (
	"blog/router"
)

//@title miniproject
//@version 1.0.0
//@description "赚圈圈API"
//@termsOfService http://swagger.io/terrms/

//@contact.name yyj
//@contact.email 2105753640@qq.com

//@host 124.221.246.5:5500
//@BasePath /api
//@Schemes http
func main() {
	router := router.Generator()

	router.Run(":8080")
}
