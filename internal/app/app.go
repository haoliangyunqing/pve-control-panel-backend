package app

import (
	"github.com/gin-gonic/gin"
	"pve-control-panel-backend/internal/api"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	router := gin.Default()
	api.RegisterRoutes(router)
	return &App{
		router: router,
	}
}
func (app *App) Run() error {
	return app.router.Run(":8080")
}
