package handlers

import (
	"../teleport"
	"log"
	"../machine"
)

type Ready struct {}

func (*Ready) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("handler---------Ready------------")

	uid := receive.From

	//用户存在
	if player,ok := machine.GLOBAL_USER[uid]; ok {
		player.Machine.Trigger( &machine.PlayerReadyState{} )
	} else {
		log.Printf("找不到用户")
	}

	return nil
}
