package greetings

import "fmt"

type User struct {
	Id   int
	Name string
}

func CreateUser(userid int, username string) *User {
	user := &User{Id: userid, Name: username}
	return user
}

// receiver
func (user *User) PrintName() {
	fmt.Println(user.Name)
}

func (user *User) ChangeName(newName string) {
	user.Name = newName
}

func (user *User) GetCredentials() (int, string) {
	return user.Id, user.Name
}

//////////////
type Puser struct {
	id   int
	name string
}

func CreatePuser(puserid int, pusername string) *Puser {
	puser := &Puser{id: puserid, name: pusername}
	return puser
}

func (puser *Puser) GetId() int {
	return puser.id
}

func (puser *Puser) GetName() string {
	return puser.name
}
