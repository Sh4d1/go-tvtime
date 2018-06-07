package tvtime

import (
	"fmt"
	"strconv"
)

// User is the struct to store a user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DisplayUser displays the current user
func DisplayUser() error {
	user, err := GetUser()
	if err != nil {
		return err
	}
	fmt.Print("Name: " + user.Name + "\n")
	fmt.Print("ID: " + strconv.Itoa(user.ID) + "\n")
	return nil
}
