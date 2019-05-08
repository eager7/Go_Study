package mcache

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMCache(t *testing.T) {
	router := gin.Default()
	cache := Initialize("localhost:8080")
	router.GET("/:who/ping", cache.MiddleWare(), PingHandle)
	if err := router.Run(); err != nil {
		panic(err)
	}
}

type Ping struct {
	who string
}

func PingHandle(context *gin.Context) {
	who := Ping{who: context.Param("who")}
	if context.BindQuery(&who) != nil {
		context.JSON(200, fmt.Sprintf("you failed, %s", who))
	}
	context.JSON(200, gin.H{"message": fmt.Sprintf("hi, %s", who.who)})
}
