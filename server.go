package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"./global"
)

func main() {
	debug.Debug = true

	//注册请求处理函数
	serverHandlers := teleport.API{
		"CreateRoom": new(handlers.CreateRoom),
		"EnterRoom": new(handlers.EnterRoom),
		"Discard": new(handlers.Discard),
		"ActionSelect": new(handlers.ActionSelect),

		teleport.HEARTBEAT : new(handlers.Heartbeat),
		teleport.IDENTITY : new(handlers.Identity),
	}

	//启动服务器
	global.SERVER.SetUID("abc").SetAPI( serverHandlers ).Server(":20125")

	select {}
}

