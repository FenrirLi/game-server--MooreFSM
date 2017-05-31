package handlers

import (
	"../teleport"
	"fmt"
	"../global"
	"../proto"
	"github.com/golang/protobuf/proto"
	"log"
)

type EnterRoom struct {}

func (*EnterRoom) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("---------enter room------------")

	// 进行解码
	request := &server_proto.EnterRoomRequest{}
	err := proto.Unmarshal(receive.Body, request)
	if err != nil {
		log.Fatal("enter room request error: ", err)
	}
	fmt.Println(request.RoomId)

	if table,ok := global.GLOBAL_TABLE[int(request.RoomId)]; ok {
		fmt.Println("enter room ",table.TableId," success")
	}else {
		//return teleport.ReturnData(nil,"CreateRoomReturn",receive.From)
	}

	return nil
}