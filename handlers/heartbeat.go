package handlers

import (
	"../teleport"
	"log"
)

type Heartbeat struct{}

func (*Heartbeat) Process(receive *teleport.NetData) *teleport.NetData {

	log.Println("-----------心跳----------",receive.From)

	return nil
}
