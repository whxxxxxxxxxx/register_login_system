package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"register_log/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
	}
	return r
}
