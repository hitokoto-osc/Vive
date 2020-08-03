package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session is a middleware func that register a session driver by Cookie
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions("ncm-session", store)
}
