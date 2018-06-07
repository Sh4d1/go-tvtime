package tvtime

import (
	"fmt"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func DisplayUser() error {
	user, err := GetUser()
	if err != nil {
		return err
	}
	fmt.Print("Name: " + user.Name + "\n")
	fmt.Print("ID: " + strconv.Itoa(user.Id) + "\n")
	return nil
}
