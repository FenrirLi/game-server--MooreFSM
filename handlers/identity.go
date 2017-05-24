package handlers

import (
	"../teleport"
	"fmt"
)

type Identity struct{}

func (*Identity) Process(receive *teleport.NetData) *teleport.NetData {
	fmt.Println("=============identity===========")
	return nil
}
