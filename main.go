package main

import (
	"ASPGo/framework"
	"ASPGo/middleware"
	"ASPGo/route"
	"net/http"
)

func main() {
	core := framework.NewCore()
	core.Use(middleware.Recovery())

	route.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
