package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	user := greetings.CreateUser(1, "Ali")
	fmt.Println(user.Id, user.Name)
	user.PrintName()
	user.ChangeName("Veli")
	user.PrintName()

	user2 := user
	user2.PrintName()
	user.ChangeName("Mehmet")
	user2.PrintName()

	id, name := user.GetCredentials()
	fmt.Println(id, name)

	puser := greetings.CreatePuser(2, "Veli")
	//fmt.Println(puser.id)
	fmt.Println(puser.GetId(), puser.GetName())
}
