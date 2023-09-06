package kernel

import (
	"ASPGo/framework"
	"net/http"
)

// 引擎服务
type HadeKernelService struct {
	engine *framework.Core
}

// 初始化web引擎服务实例
func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*framework.Core)
	return &HadeKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
