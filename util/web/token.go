package web

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ValidTokenByContext is a func that valid the authorization by bearer token
func ValidTokenByContext(ctx *gin.Context) (ok bool) {
	ok = false
	masterKey := viper.GetString("server.auth.master_key")
	token, isOk := ParseBearerTokenFromHeader(ctx)
	if !isOk && token == masterKey { // fallback: support token that not encoded by base64
		ok = true
	} else if token == masterKey {
		ok = true
	}
	return
}
