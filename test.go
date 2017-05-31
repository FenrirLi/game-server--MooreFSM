package main

import (
	"./proto"
	"log"
	"github.com/golang/protobuf/proto"
)

func main() {
	create_req := &create_room.CreateRoomRequest{"1111",2}

	// 进行编码
	data, err := proto.Marshal(create_req)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	log.Println(data)

	// 进行解码
	new_req := &create_room.CreateRoomRequest{}
	err = proto.Unmarshal(data, new_req)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("uid:%s    round:%d;",new_req.Uuid,new_req.Round)

	// 测试结果
	if create_req.String() != new_req.String() {
		log.Fatalf("data mismatch %q != %q", create_req.String(), new_req.String())
	}
}
