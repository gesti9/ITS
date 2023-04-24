package service

import "fmt"

type User struct {
	name    string
	pasword string
}

func UserList(log string, pas string) bool {
	admin := User{"sultan", "12345"}
	if log == admin.name && pas == admin.pasword {
		fmt.Println("Hello admin")
		return true
	}

	return false
}
