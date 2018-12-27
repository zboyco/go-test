package main

import (
	"go-test/protobuf/models"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("测试protobuf")
	// 创建对象
	group := models.UserGroup{
		Id:   1,
		Name: "测试组",
	}

	for i := 0; i < 3; i++ {

		m := models.User{
			Id:   int32(i),
			Name: "Test",
			Pwd:  "Test",
		}
		m.Phones = append(m.Phones, "1")
		m.Phones = append(m.Phones, "2")
		m.Phones = append(m.Phones, "3")
		group.Users = append(group.Users, &m)
	}

	// 编码
	mData, err := proto.Marshal(&group)
	if err != nil {
		log.Fatalln(err)
	}

	// 解码
	var obj models.UserGroup
	err = proto.Unmarshal(mData, &obj)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(mData)

	// 输出
	log.Println(obj.Id, obj.Name)
	for _, v := range obj.Users {
		log.Println("UserID:", v.Id)
		log.Println("UserName:", v.Name)
		for _, p := range v.Phones {
			log.Println("Phone:", p)
		}
	}
}
