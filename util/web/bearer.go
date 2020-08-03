package web

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

// ParseBearerTokenFromHeader is a func that parse bearer token in authorization header
func ParseBearerTokenFromHeader(ctx *gin.Context) (token string, ok bool) {
	authorization := ctx.GetHeader("Authorization")
	// parse authorization
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		ok = false
		return
	}
	token = strings.ReplaceAll(authorization, "Bearer ", "")
	// base64 decode
	buffer, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Error(err)
		ok = false
		return
	}
	ok = true
	token = string(buffer[:])
	return
}
