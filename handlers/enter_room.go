package handlers

import (
	"../teleport"
	"log"
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

	//房间存在
	if table,ok := machine.GLOBAL_TABLE[int(request.RoomId)]; ok {
		log.Println("enter room ",table.TableId," success")

		//创建玩家,依据桌子房间（桌子）情况分配座位
		player := machine.CreatePlayer( uid, table )
		//创建玩家的状态机
		player_machine := machine.NewPlayerMachine( &player, &machine.PlayerInitState{}, nil )
		player.Machine = &player_machine
		//记录玩家信息到桌子
		table.PlayerDict[player.Seat] = &player
		//记录全局用户
		machine.GLOBAL_USER[player.Uid] = &player

		for key, value := range table.PlayerDict {
			log.Println("Key:", key, "Value:", value.Uid)
		}

		//是否全部准备
		//if table.IsAllReady() {
		//	log.Println("all player are ready for game")
		//	table.Machine.Trigger( &machine.TableReadyState{} )
		//}

	} else {
		//return teleport.ReturnData(nil,"CreateRoomReturn",receive.From)
	}

	return nil
}