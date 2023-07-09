package framework

import (
	"net/http"
	"strings"
)

type Core struct {
	router map[string]map[string]ControllerHandler
}

func NewCore() *Core {
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{}
}

func (c *Core) Get(url string, hanler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = hanler
}

func (c *Core) Post(url string, hanler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = hanler
}

func (c *Core) Put(url string, hanler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = hanler
}

func (c *Core) Delete(url string, hanler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = hanler
}

func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.RawPath
	method := request.Method

	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	var mehtodHandlers map[string]ControllerHandler
	var ok bool
	if mehtodHandlers, ok = c.router[upperUri]; !ok {
		return nil
	}

	var handler ControllerHandler
	if handler, ok = mehtodHandlers[upperMethod]; !ok {
		return nil
	}

	return handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) { // TODO
	ctx := NewContext(request, response)

	router := c.FindRouterByRequest(request)
	if router == nil {
		// 如果没有找到，这里打印日志
		ctx.Json(404, "not found")
		return
	}

	// 调用路由函数，如果返回err 代表存在内部错误，返回500状态码
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
