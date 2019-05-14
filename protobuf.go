package main

import (
	"github.com/golang/protobuf/proto"

	"fmt"
	"ind/pb"
)

func main() {
	//初始化message UserInfo
	usermsg := &pb.UserInfo{
		UserType: 1,
		UserName: "Jok",
		UserInfo: "I am a woker!",
	}

	//序列化
	userdata, err := proto.Marshal(usermsg)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	}

	//反序列化
	encodingmsg := &pb.UserInfo{}
	err = proto.Unmarshal(userdata, encodingmsg)

	if err != nil {
		fmt.Println("Unmarshaling error: ", err)
	}

	fmt.Printf("GetUserType: %d\n", encodingmsg.GetUserType())
	fmt.Printf("GetUserName: %s\n", encodingmsg.GetUserName())
	fmt.Printf("GetUserInfo: %s\n", encodingmsg.GetUserInfo())
}
