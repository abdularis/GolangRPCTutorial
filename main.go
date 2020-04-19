package main

import (
	"fmt"
	"golangrpctutorial/model"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

func save() {
	user := model.User{
		Id:   12,
		Name: "Abdul Aris",
	}

	data, err := proto.Marshal(&user)
	if err != nil {
		fmt.Println("message marshalling failed")
		return
	}

	err = ioutil.WriteFile("./user.out", data, 0644)
	if err != nil {
		fmt.Println("failed to write to file")
	}
}

func read() {
	data, err := ioutil.ReadFile("./user.out")
	if err != nil {
		fmt.Println("error read file")
		return
	}

	var user model.User
	err = proto.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("error marshalling data to User model")
		return
	}

	fmt.Println(user.String())
}

func main() {
	save()
	read()
}
