package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"log"
	//"time"
	"github.com/golang/protobuf/proto"
	"./proto"
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var table_id string

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	uid := "C1"

	//注册请求处理函数
	clientHandlers := teleport.API{
		"CreateRoomReturn" : new(ClientHeartBeat),
		"DiscardResponse" : new(DiscardResponse),
		"DrawResponse" : new(DrawResponse),
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

	//指令
	running := true
	var inp []byte
	var data2 []byte
	var order string
	var card int
	var req2 = &server_proto.DiscardRequest{}
	for running {
		fmt.Println("please input :")
		reader := bufio.NewReader(os.Stdin)
		inp, _, _ = reader.ReadLine()
		order = string(inp)
		if order == "stop"{
			running = false
		} else if order == "discard" {
			inp, _, _ = reader.ReadLine()
			order = string(inp)
			card, _ = strconv.Atoi(order)
			req2.Card = int32(card)
			data2, _ = proto.Marshal(req2)
			tp.Request(data2, "Discard", "discard_flag")
			log.Println("discard ",card)
		}
	}

	select {}
}

type ClientHeartBeat struct{}
func (*ClientHeartBeat) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============room create return===============")

	// 进行解码
	response := &server_proto.CreateRoomResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(response.RoomId)

	request := &server_proto.EnterRoomRequest{response.RoomId}
	data := server_proto.MessageEncode(request)

	return teleport.ReturnData( data, "EnterRoom" )
}

type DrawResponse struct{}
func (*DrawResponse) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============DrawResponse===============")

	// 进行解码
	response := &server_proto.DrawCardResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(" draw ",response.Card)
	return nil
}

type DiscardResponse struct{}
func (*DiscardResponse) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============DiscardResponse===============")

	// 进行解码
	response := &server_proto.DiscardResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(response.Uuid," discard ",response.Card)
	return nil
}