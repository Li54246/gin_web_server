package system

import (
	"gin_web_server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (t *TestApi) Test(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}
