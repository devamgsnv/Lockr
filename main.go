package main

import (
	"fmt"
)

func main() {
	login := promptData("Enter login")
	password := promptData("Enter password")
	url := promptData("Enter URL")
	accountGoogle, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Incorrect format URL or Login")
		return
	}
	accountGoogle.outputPassword()
	fmt.Println(accountGoogle)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
