package server

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *Application) InitRoutes() {
	var r = app.Router
	internal := r.Group("/internal")
	{
		internal.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		internal.GET("/raw-db", func(c *gin.Context) {
			c.File(app.Cfg.FilePath)
		})
	}
	rootRouter := r.Group("/api")
	addDynamicRoutes(rootRouter, app)
}

func addDynamicRoutes(router *gin.RouterGroup, app *Application) {
	data := app.Database.Data
	for k := range *data {
		addGetAll(k, router, app)
		getById(fmt.Sprintf("%s/:id", k), k, router, app)
		addToFile(k,router,app)
	}
}

func addToFile(prefix string, router *gin.RouterGroup, app *Application) {
	router.POST(prefix, func (c *gin.Context) {
		var json interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Added to %s", prefix)})
		
		//Refactor this
		data := *app.Database.Data
		r  := data[prefix]
		var kindOfResource = reflect.TypeOf(r).Kind()
		if kindOfResource == reflect.Array || kindOfResource == reflect.Slice {
		dataArr := r.([]interface{})
		dataArr = append(dataArr, json)
		data[prefix] = dataArr
		if !app.Cfg.ReadOnlyMode {
			// Implement writing to file.
		}
	}
	})
}

func addGetAll(prefix string, router *gin.RouterGroup, app *Application) {
	data := *app.Database.Data
	router.GET(prefix, func(c *gin.Context) {
		c.JSON(200, gin.H{
			prefix: data[prefix],
		})
	})
}

func getById(prefix string, resource string, router *gin.RouterGroup, app *Application) {
	data := *app.Database.Data
	r  := data[resource]
	var kindOfResource = reflect.TypeOf(r).Kind()
	if kindOfResource == reflect.Array || kindOfResource == reflect.Slice {
		dataArr := r.([]interface{})
		router.GET(prefix, func(c *gin.Context) {
			id := c.Param("id")
			intId,err := strconv.Atoi(id)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
				return
			}
			if intId-1 > len(dataArr)-1 || intId-1 < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": intId})
				return
			}
			c.JSON(200, gin.H{
				id: dataArr[intId-1],
			})
		})
	}
}
