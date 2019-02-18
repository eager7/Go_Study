package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BaseResp struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type GetActionTypesRsp struct {
	BaseResp
	Data TypesData `json:"data"`
}

type TypesData struct {
	Types []string `json:"types"`
}

type Ping struct {
	who string
}

func InitializeGin() {
	router := gin.Default()
	router.GET("/:who/ping", PingHandle)
	if err := router.Run(); err != nil {
		panic(err)
	}
}

func PingHandle(context *gin.Context) {
	who := Ping{who: context.Param("who")}
	if context.BindQuery(&who) != nil {
		resp := GetActionTypesRsp{
			BaseResp: BaseResp{
				Errno:  1,
				Errmsg: "bind failed",
			},
			Data: TypesData{
				Types: nil,
			},
		}
		context.JSON(200, resp)
	}
	context.JSON(200, gin.H{"message": fmt.Sprintf("%s pong", who.who)})
}
