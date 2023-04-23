package system

import (
	v1 "gin_web_server/api/v1"
	"github.com/gin-gonic/gin"
)

type TestRouter struct {
}

func (s *TestRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	testRouter := Router.Group("test")
	TestApi := v1.ApiGroupApp.SystemApiGroup.TestApi
	{
		testRouter.GET("login", TestApi.Test)
	}
	return testRouter
}
