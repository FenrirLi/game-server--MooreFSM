package handlers

import (
	"../teleport"
	"fmt"
)

type Heartbeat struct{}

func (*Heartbeat) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("-----------心跳----------")
	fmt.Println(receive.Body)
	fmt.Println(receive.Operation)
	fmt.Println(receive.From)
	fmt.Println(receive.To)
	fmt.Println(receive.Status)
	fmt.Println(receive.Flag)

	return teleport.ReturnData("confirm heartbeat",teleport.HEARTBEAT,receive.From)
}
