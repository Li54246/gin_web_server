package router

import "gin_web_server/router/system"

type RouterGroup struct {
	System system.RouterGroup
	//Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
