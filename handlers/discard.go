package handlers

import (
	"../teleport"
	"log"
	"../machine"
)

type Discard struct {}

func (*Discard) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("handler---------discard------------")

	uid := receive.From

	//// 进行解码
	//request := &server_proto.DiscardRequest{}
	//err := proto.Unmarshal(receive.Body, request)
	//if err != nil {
	//	log.Fatal("discard request error: ", err)
	//}

	//用户存在
	if player,ok := machine.GLOBAL_USER[uid]; ok {
		player.Machine.CurrentState.Execute( player, machine.PlayerEvent["PLAYER_EVENT_DISCARD"], receive.Body )
	} else {
		log.Printf("找不到用户")
	}

	return nil
}