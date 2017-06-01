package handlers

import (
	"../teleport"
	"log"
	"../global"
	"../proto"
	"github.com/golang/protobuf/proto"
	"../machine"
)

type EnterRoom struct {}

func (*EnterRoom) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("---------enter room------------")

	uid := receive.From

	// 进行解码
	request := &server_proto.EnterRoomRequest{}
	err := proto.Unmarshal(receive.Body, request)
	if err != nil {
		log.Fatal("enter room request error: ", err)
	}

	if table,ok := global.GLOBAL_TABLE[int(request.RoomId)]; ok {
		log.Println("enter room ",table.TableId," success")

		player := machine.CreatePlayer( uid, table )
		player_machine := machine.NewPlayerMachine( player, machine.PlayerEventStatus["EVENT_READY"], nil )
		player.Machine = &player_machine

		table.PlayerDict[player.Seat] = player

		for key, value := range table.PlayerDict {
			log.Println("Key:", key, "Value:", value.Uid)
		}

		if table.IsAllReady() {
			log.Println("all player are ready for game")
			table.Machine.Trigger( &machine.TableReadyStatus{} )
		}

	}else {
		//return teleport.ReturnData(nil,"CreateRoomReturn",receive.From)
	}

	return nil
}