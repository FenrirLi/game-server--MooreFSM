package handlers

import (
	"../teleport"
	"fmt"
)

type CreateRoom struct {}

func (*CreateRoom) Process(receive *teleport.NetData) *teleport.NetData {

	fmt.Println("---------------------")
	fmt.Println(receive.Body)
	fmt.Println(receive.Operation)
	fmt.Println(receive.From)
	fmt.Println(receive.To)
	fmt.Println(receive.Status)
	fmt.Println(receive.Flag)

	return nil
}