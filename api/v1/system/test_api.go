package system

import (
	"gin_web_server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (t *TestApi) Test(c *gin.Context) {
	num, _ := serviceTestApi.SetTest(10)
	response.OkWithDetailed(num, "登录成功", c)
}
