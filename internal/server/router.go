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
	if app.Cfg.StaticFiles != "" || !(len(app.Cfg.StaticFiles) > 0) {
		r.Static("/assets", app.Cfg.StaticFiles)
	}
	rootRouter := r.Group("/api")
	addDynamicRoutes(rootRouter, app)
}

func addDynamicRoutes(router *gin.RouterGroup, app *Application) {
	data := app.Database.Data
	for k := range *data {
		addGetAll(k, router, app)
		getById(fmt.Sprintf("%s/:id", k), k, router, app)
		addToDb(k, router, app)
		deleteFromDb(fmt.Sprintf("%s/:id", k), k, router, app)
	}
}

func deleteFromDb(prefix string, resource string, router *gin.RouterGroup, app *Application) {
	data := *app.Database.Data
	r := data[resource]
	var kindOfResource = reflect.TypeOf(r).Kind()
	if kindOfResource == reflect.Array || kindOfResource == reflect.Slice {
		dataArr := r.([]interface{})
		router.DELETE(prefix, func(c *gin.Context) {
			id := c.Param("id")
			intId, err := strconv.Atoi(id)
			if err != nil || intId-1 > len(dataArr)-1 || intId-1 < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
				return
			}
			dataArr := r.([]interface{})
			dataArr = append(dataArr[:intId-1], dataArr[intId+1:]...)
			fmt.Println(dataArr...)
			data[resource] = dataArr
			c.JSON(http.StatusAccepted, gin.H{
				id: id,
			})
			if !app.Cfg.ReadOnlyMode {
				app.Database.WriteToDB(app.Cfg.FilePath)
			}
		})
	}
}

func addToDb(prefix string, router *gin.RouterGroup, app *Application) {
	router.POST(prefix, func(c *gin.Context) {
		var json interface{}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data := *app.Database.Data
		r := data[prefix]
		var kindOfResource = reflect.TypeOf(r).Kind()
		if kindOfResource == reflect.Array || kindOfResource == reflect.Slice {
			dataArr := r.([]interface{})
			dataArr = append(dataArr, json)
			data[prefix] = dataArr
			c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Added to %s", prefix)})
			if !app.Cfg.ReadOnlyMode {
				app.Database.WriteToDB(app.Cfg.FilePath)
			}
		} else {
			data[prefix] = json
			c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Added to %s", prefix)})
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
	r := data[resource]
	var kindOfResource = reflect.TypeOf(r).Kind()
	if kindOfResource == reflect.Array || kindOfResource == reflect.Slice {
		dataArr := r.([]interface{})
		router.GET(prefix, func(c *gin.Context) {
			id := c.Param("id")
			intId, err := strconv.Atoi(id)
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
