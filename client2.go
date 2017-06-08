package main

import (
	"./teleport"
	"./teleport/debug"
	"./handlers"
	"./proto"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"github.com/golang/protobuf/proto"
	"log"
)

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	uid := "C2"
	room_id := int32(111111)

	//注册请求处理函数
	clientHandlers := teleport.API{
		"DiscardResponse" : new(DiscardResponse2),
		"DrawResponse" : new(DrawResponse2),
		teleport.IDENTITY : new(handlers.Identity),
	}

	//启动客户端
	tp := teleport.New().SetUID(uid, "abc").SetAPI( clientHandlers )
	tp.Client("127.0.0.1", ":20125")

	request := &server_proto.EnterRoomRequest{room_id}
	data := server_proto.MessageEncode(request)

	tp.Request(data, "EnterRoom", "enter_room_flag")

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

type DrawResponse2 struct{}
func (*DrawResponse2) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============DrawResponse===============")

	// 进行解码
	response := &server_proto.DrawCardResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(" draw ",response.Card)
	return nil
}

type DiscardResponse2 struct{}
func (*DiscardResponse2) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============DiscardResponse===============")

	// 进行解码
	response := &server_proto.DiscardResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(response.Uuid," discard ",response.Card)
	return nil
}
