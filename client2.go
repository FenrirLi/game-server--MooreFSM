package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"./proto"
)

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	uid := "C2"
	room_id := int32(111111)

	//注册请求处理函数
	clientHandlers := teleport.API{
		teleport.IDENTITY : new(handlers.Identity),
	}

	//启动客户端
	tp := teleport.New().SetUID(uid, "abc").SetAPI( clientHandlers )
	tp.Client("127.0.0.1", ":20125")

	request := &server_proto.EnterRoomRequest{room_id}
	data := server_proto.MessageEncode(request)

	tp.Request(data, "EnterRoom", "enter_room_flag")
	select {}
}
