package main

import (
	"fmt"

	"example.com/models"
)

func main() {
	user := models.CreateUser(1, "Ali")
	fmt.Println(user.Id, user.Name)
	user.PrintName()
	user.ChangeName("Veli")
	user.PrintName()

	user2 := user
	user2.PrintName()
	user.ChangeName("Mehmet")
	user2.PrintName()

	user3 := *user
	user3.PrintName()
	user.ChangeName("Emin")
	user3.PrintName()

	id, name := user.GetCredentials()
	fmt.Println(id, name)

	puser := models.CreatePuser(2, "Veli")
	//fmt.Println(puser.id)
	fmt.Println(puser.GetId(), puser.GetName())
}
