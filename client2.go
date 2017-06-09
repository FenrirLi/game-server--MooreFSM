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
	"log"
)

// 有标识符UID的demo，保证了客户端链接唯一性
func main() {

	debug.Debug = true

	uid := "C2"
	room_id := int32(111111)

	//注册请求处理函数
	clientHandlers := teleport.API{
		"ActionResponse" : new(ActionResponse2),
		"ActionPrompt" : new(ActionPrompt2),
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
	for running {
		fmt.Println("please input :")
		reader := bufio.NewReader(os.Stdin)
		inp, _, _ = reader.ReadLine()
		order = string(inp)
		if order == "stop"{
			running = false
		} else if order == "discard" {
			inp, _, _ = reader.ReadLine()
			card_input := string(inp)
			card, _ := strconv.Atoi(card_input)
			request := &server_proto.DiscardRequest{
				int32(card),
			}
			data2 = server_proto.MessageEncode( request )
			tp.Request(data2, "Discard", "discard_flag")
		} else if order == "action" {
			inp, _, _ = reader.ReadLine()
			select_id_input := string(inp)
			select_id, _ := strconv.Atoi(select_id_input)
			request := &server_proto.ActionSelectRequest{
				int32(select_id),
			}
			data2 = server_proto.MessageEncode( request )
			tp.Request(data2, "ActionSelect", "action_select")
		} else if order == "ready" {
			tp.Request(nil, "Ready", "ready_flag")
		}
	}

	select {}
}

type ActionResponse2 struct{}
func (*ActionResponse2) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============ActionResponse===============")

	// 进行解码
	response := &server_proto.ActionResponse{}
	server_proto.MessageDecode( receive.Body, response )
	log.Println(response.Uuid," ",response.ActionName," ",response.Card)
	return nil
}

type ActionPrompt2 struct {}
func (*ActionPrompt2) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("=============ActionPrompt===============")

	// 进行解码
	response := &server_proto.ActionPrompt{}
	server_proto.MessageDecode( receive.Body, response )
	for _,v := range response.Action {
		log.Println( v )
	}
	return nil
}