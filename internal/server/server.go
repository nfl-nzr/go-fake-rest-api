package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nfl-nzr/go-fake-rest-api/internal/db"
)

type Config struct {
	Port         int
	FilePath     string
	ReadOnlyMode bool
	Env          string
}

type Application struct {
	Cfg      Config
	Router   *gin.Engine
	Database db.Database
}

func (a *Application) CreateServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	a.Router = r
	a.InitDb()
	a.InitRoutes()
}

func (a *Application) InitDb() error {
	// Handle db connections here
	if err := a.Database.Connect(a.Cfg.FilePath); err != nil {
		return err
	}
	return nil
}

func (a *Application) StartServer() error {
	return a.Router.Run(fmt.Sprintf(":%d", a.Cfg.Port))
}
