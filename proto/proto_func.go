package server_proto

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func MessageEncode( request proto.Message ) []byte {
	data, err := proto.Marshal( request )
	if err != nil {
		log.Fatal("message encode error: ", err)
	}
	return data
}

func MessageDecode( data []byte, message proto.Message ) {
	// 进行解码
	err := proto.Unmarshal( data, message )
	if err != nil {
		log.Fatal("create room request error: ", err)
	}
}