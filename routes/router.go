package routes

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/hitokoto-osc/Vive/controllers/v1"
	"github.com/hitokoto-osc/Vive/middlewares"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setupRoutes(r *gin.Engine) {
	if !viper.IsSet("server.secret") {
		log.Fatal("[web] can't start server because of the secret is not set.")
	}
	r.Use(middlewares.Session(viper.GetString("server.secret")))

	// Setup router
	r.GET("/", func(context *gin.Context) {
		context.String(200, "Hello, World.")
	})

	v1 := r.Group("/v1")
	{
		// protected routes
		// protected := r.Group("/api/v1", middlewares.AuthByMasterKey())
		// {
		// }
		// common routes
		v1.GET("/ping", apiV1.Ping)
	}
}
