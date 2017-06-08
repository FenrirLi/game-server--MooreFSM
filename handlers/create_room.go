package handlers

import (
	"../teleport"
	"log"
	"../machine"
	"../proto"
)

type CreateRoom struct {}

func (*CreateRoom) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("---------create room------------")

	// 进行解码
	request := &server_proto.CreateRoomRequest{}
	server_proto.MessageDecode( receive.Body, request )
	log.Println(request.Uuid)
	log.Println(request.Round)


	table := machine.CreateTable( receive.From )
	new_machine := machine.NewTableMachine( &table, nil, nil )
	table.Machine = &new_machine
	machine.GLOBAL_TABLE[table.TableId] = &table

	//返回
	response := &server_proto.CreateRoomResponse{int32(table.TableId)}
	data := server_proto.MessageEncode( response )

	return teleport.ReturnData(data,"CreateRoomReturn")
}