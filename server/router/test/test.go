package test

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (e *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.TestApiGroup.TestApi
	{
		testRouter.GET("t1", testApi.Test1)
	}
}
