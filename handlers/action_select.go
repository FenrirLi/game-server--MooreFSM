package handlers

import (
	"../teleport"
	"log"
	"../machine"
)

type ActionSelect struct {}

func (*ActionSelect) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("handler---------action select------------")

	uid := receive.From

	//用户存在
	if player,ok := machine.GLOBAL_USER[uid]; ok {
		player.Machine.CurrentState.Execute( player, machine.PlayerEvent["PLAYER_EVENT_ACTION"], receive.Body )
	} else {
		log.Printf("找不到用户")
	}

	return nil
}