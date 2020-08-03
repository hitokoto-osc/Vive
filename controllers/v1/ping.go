package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hitokoto-osc/Vive/util/web"
)

// Ping is a controller func that impl a Pong response,
// intended to notify the client that server is ok
func Ping(c *gin.Context) {
	web.Success(c, map[string]interface{}{})
}
