package main

import (
	"Lockr/account"
	"Lockr/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promptData("Enter login")
	password := promptData("Enter password")
	url := promptData("Enter URL")
	accountGoogle, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Incorrect format URL or Login")
		return
	}
	file, err := accountGoogle.ToBytes()
	if err != nil {
		fmt.Println("Data could not be converted in JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
