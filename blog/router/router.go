package router

import (
	"blog/handler/getinfor"
	"blog/handler/star"
	"blog/handler/user"
	"blog/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//hello
func Generator() *gin.Engine {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", getinfor.Res)
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	post := router.Group("/api/post")
	{
		post.POST("/person/register", user.Register)
		post.POST("/person/login", user.Login)
	}

	get := router.Group("/api/get")
	{
		get.GET("/file/data", getinfor.GetBlog)
		get.GET("/file/name", getinfor.GetName)
	}

	star_group := router.Group("/api/star")
	star_group.Use(middleware.Parse)
	{
		star_group.GET("/delete", star.DeleteStar)
		star_group.GET("/addition", star.AddStar)
	}

	return router
}
