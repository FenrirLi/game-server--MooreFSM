package handlers

import (
	"../teleport"
	"log"
)

type Identity struct{}

func (*Identity) Process(receive *teleport.NetData) *teleport.NetData {
	log.Println("=============identity===========")
	return nil
}
