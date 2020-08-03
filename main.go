package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Vive/config"
	"github.com/hitokoto-osc/Vive/flag"
	"github.com/hitokoto-osc/Vive/prestart"
	"github.com/hitokoto-osc/Vive/routes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"runtime"
)

var (
	// BuildTag is a commit hash that will be injected in release mode
	BuildTag = "Unknown"
	// BuildTime is a time, when it build, that will be injected in release mode
	BuildTime = "Unknown"
	// MakeVersion is the version of make in release, will be injected in release mode
	MakeVersion = "Unknown"
	// Version is the version of this program, will be injected in release mode
	Version = "development"
)

var r *gin.Engine

func init() {
	// global set build information
	config.BuildTag = BuildTag
	config.BuildTime = BuildTime
	config.GoVersion = runtime.Version()
	config.MakeVersion = MakeVersion
	config.Version = Version

	// Parse Flag
	flag.Parse()

	// Init Drivers
	prestart.Do()

	if config.Debug {
		log.Info("[debug] 已启用调试模式")
	}

	// init Web Server
	r = routes.InitWebServer()
}

func main() {
	// start Server
	r.Run(":" + viper.GetString("server.port"))
}
