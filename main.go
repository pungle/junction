//Copyright (C) Mr.Pungle

package main

import (
	"github.com/pungle/dawn/web"
	"time"
)

func HelloHandler(ctx *web.HttpContext) {
	ctx.Response.Write([]byte("hello..."))
}

func main() {
	config := web.NewConfig(":80", web.DEFAULT_LOG_FLAG, web.DEFAULT_LOG_LEVEL, false, "", "")
	driver := web.NewRedisSessionDriver("tcp", ":6379", 100, 1000, 30*time.Second)
	sessionCtx := web.NewSessionContext(
		driver,
		"sid",
		"junction.cn",
		24*365*time.Hour,
		"/",
		true,
		false,
		7*24*time.Hour,
	)
	server := web.NewServer(config, sessionCtx, nil)
	server.AddHandler("/", HelloHandler)
	println(server.Start().Error())
	println("come?")
}
