package main

import (
	"ASPGo/framework"
	"ASPGo/route"
	"net/http"
)

func main() {
	core := framework.NewCore()
	route.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}
