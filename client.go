package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"fmt"
	//"time"
	"github.com/golang/protobuf/proto"
	"log"
	"./proto"
)

var table_id string

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	uid := "C1"

	//注册请求处理函数
	clientHandlers := teleport.API{
		"CreateRoomReturn" : new(ClientHeartBeat),
		//teleport.HEARTBEAT : new(ClientHeartBeat),
		teleport.IDENTITY : new(handlers.Identity),
	}

	//启动客户端
	tp := teleport.New().SetUID(uid, "abc").SetAPI( clientHandlers )
	tp.Client("127.0.0.1", ":20125")

	req := &server_proto.CreateRoomRequest{uid,4}
	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatal("create room request error: ", err)
	}

	tp.Request(data, "CreateRoom", "create_room_flag")
	//time.Sleep(5*time.Second)
	//tp.Request("111111", "EnterRoom", "enter_room_flag")
	select {}
}

type ClientHeartBeat struct{}
func (*ClientHeartBeat) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("=============room create return===============")

	// 进行解码
	response := &server_proto.CreateRoomResponse{}
	err := proto.Unmarshal(receive.Body, response)
	if err != nil {
		log.Fatal("create room response error: ", err)
	}
	fmt.Println(response.RoomId)

	request := &server_proto.EnterRoomRequest{response.RoomId}
	data, err := proto.Marshal( request )
	if err != nil {
		log.Fatal("enter room request error: ", err)
	}

	return teleport.ReturnData( data, "EnterRoom" )
}
