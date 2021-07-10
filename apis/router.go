package apis

import (
	"github.com/gin-gonic/gin"
	apis "hongbao/apis/api/envelope"
)

func APIRouter() *gin.Engine {
	r:=gin.Default()


	v1 := r.Group("/v1")
	{
		v1.GET("/login", apis.Hi)

		//创建红包
		v1.POST("/CreateEnvelope", apis.CreateEnvelope)

	}


	return r
}

