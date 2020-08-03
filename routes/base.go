package routes

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Vive/config"
	"github.com/hitokoto-osc/Vive/middlewares"
	"github.com/thinkerou/favicon"
	"os"
)

// InitWebServer is a web server register, implemented by gin
func InitWebServer() *gin.Engine {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// load middleware
	r.Use(requestid.New())

	// favicon.ico
	if _, err := os.Stat("../static/resource/favicon.ico"); err == nil {
		r.Use(favicon.New("../static/resource/favicon.ico"))
	} else if _, err := os.Stat("./static/resource/favicon.ico"); err == nil {
		r.Use(favicon.New("./static/resource/favicon.ico"))
	}

	r.Use(middlewares.Cors())

	// setup routes
	setupRoutes(r)
	return r
}
