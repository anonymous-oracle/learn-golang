package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
	fmt.Println(bs)
	loginPwd := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("Unauthorised")
		return
	}
	fmt.Println("You're logged in")
}
