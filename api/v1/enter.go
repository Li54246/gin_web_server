package v1

import "gin_web_server/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	//ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
