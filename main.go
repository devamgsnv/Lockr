package main

import (
	"Lockr/account"
	"Lockr/files"
	"fmt"
)

func main() {
	files.WriteFile("Hello!", "file.txt")
	login := promptData("Enter login")
	password := promptData("Enter password")
	url := promptData("Enter URL")
	accountGoogle, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Incorrect format URL or Login")
		return
	}
	accountGoogle.OutputPassword()
	fmt.Println(accountGoogle)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
