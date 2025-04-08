package main

import (
	"Lockr/account"
	"fmt"
)

func main() {
	fmt.Println("Lockr")
Menu:
	for {
		option := getMenu()
		switch option {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
	createAccount()
}

func getMenu() int {
	var option int
	fmt.Println("Choice option")
	fmt.Println("1. Create account")
	fmt.Println("2. Search account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Scan(&option)
	return option
}

func findAccount() {

}

func deleteAccount() {

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
	vault := account.NewVault()
	vault.AddAccount(*accountGoogle)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
