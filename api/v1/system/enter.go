package system

import "gin_web_server/service"

type ApiGroup struct {
	TestApi
}

var (
	serviceTestApi = service.ServiceGroupApp.SystemServiceGroup
)
