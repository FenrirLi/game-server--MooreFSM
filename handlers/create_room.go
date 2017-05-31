package handlers

import (
	"../teleport"
	"fmt"
	"../machine"
	"../global"
	"../proto"
	"github.com/golang/protobuf/proto"
	"log"
)

type CreateRoom struct {}

func (*CreateRoom) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("---------create room------------")

	// 进行解码
	request := &server_proto.CreateRoomRequest{}
	err := proto.Unmarshal(receive.Body, request)
	if err != nil {
		log.Fatal("create room request error: ", err)
	}
	fmt.Println(request.Uuid)
	fmt.Println(request.Round)


	table := machine.CreateTable( receive.From )
	new_machine := machine.NewTableMachine( table, nil, nil )
	table.Machine = &new_machine
	global.GLOBAL_TABLE[table.TableId] = &table

	response := &server_proto.CreateRoomResponse{int32(table.TableId)}
	data, err := proto.Marshal(response)
	if err != nil {
		log.Fatal("create room response error: ", err)
	}

	return teleport.ReturnData(data,"CreateRoomReturn")
}