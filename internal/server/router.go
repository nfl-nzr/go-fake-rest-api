package server

import "github.com/gin-gonic/gin"

func (a *Application) InitRoutes() {
	var r = a.Router
	internal := r.Group("/internal")
	{
		internal.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		internal.GET("/raw-db", func (c *gin.Context)  {
			c.File(a.Cfg.FilePath)
		})
	}
	root := r.Group("/") 
	{
		root.GET("/")
	}
}
