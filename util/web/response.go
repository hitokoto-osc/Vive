package web

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"time"
)

var errorMessageMap = map[int]string{
	401: "Unauthorized.",
	404: "Not found specific route/file.",
	403: "Forbidden.",
	400: "Bad request.",
	500: "Server error. If occurs frequently, please contact the author.",
}

// Success is a func that do the common situation of responding successful formation
func Success(ctx *gin.Context, data map[string]interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code":       200,
		"message":    "ok",
		"data":       data,
		"request_id": requestid.Get(ctx),
		"ts":         time.Now().UnixNano() / 1e6,
	})
}

// Fail is a func that do the common situation of responding failed formation
func Fail(ctx *gin.Context, data map[string]interface{}, code int) {
	var status int
	if code <= 0 {
		status = 200
	} else {
		status = code
	}
	msg, ok := errorMessageMap[code]
	if !ok {
		msg = "Unknown status code, please contact author."
	}
	ctx.JSON(status, map[string]interface{}{
		"code":       code,
		"message":    msg,
		"data":       data,
		"request_id": requestid.Get(ctx),
		"ts":         time.Now().UnixNano() / 1e6,
	})
}
