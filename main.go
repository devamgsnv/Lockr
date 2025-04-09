package main

import (
	"Lockr/account"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Lockr")
	vault := account.NewVault()
Menu:
	for {
		option := getMenu()
		switch option {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
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

func findAccount(vault *account.Vault) {
	url := promptData("Enter URL for search")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("No accounts found")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Enter URL for search")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Deleted")
	} else {
		color.Red("Don't have")
	}
}

func createAccount(vault *account.Vault) {
	login := promptData("Enter login")
	password := promptData("Enter password")
	url := promptData("Enter URL")
	accountGoogle, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Incorrect format URL or Login")
		return
	}

	vault.AddAccount(*accountGoogle)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
