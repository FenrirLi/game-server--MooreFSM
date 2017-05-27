package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"fmt"
	//"time"
)

var table_id string

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	//注册请求处理函数
	clientHandlers := teleport.API{
		"CreateRoomReturn" : new(ClientHeartBeat),
		//teleport.HEARTBEAT : new(ClientHeartBeat),
		teleport.IDENTITY : new(handlers.Identity),
	}

	//启动客户端
	tp := teleport.New().SetUID("C1", "abc").SetAPI( clientHandlers )
	tp.Client("127.0.0.1", ":20125")

	//tp.Request("客户端请求创建房间", "CreateRoom", "create_room_flag")
	//time.Sleep(5*time.Second)
	tp.Request("111111", "EnterRoom", "enter_room_flag")
	select {}
}

type ClientHeartBeat struct{}
func (*ClientHeartBeat) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("=============room create return===============")
	fmt.Println(receive.Body)
	fmt.Println(receive.Operation)
	fmt.Println(receive.From)
	fmt.Println(receive.To)
	fmt.Println(receive.Status)
	fmt.Println(receive.Flag)

	table_id = receive.Body

	return nil
}
