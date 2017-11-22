package main

import (
	"fmt"
	"goFile/utils"
	"os"
)

var usersChoice string
var phoneNumber string
var name string

func main() {

	for {
		menu()
		if usersChoice == "0" {
			fmt.Println("До свидания.")
			break
		}
	}
}

func menu() {
	printMenuAndSeparator()
	fmt.Scanln(&usersChoice)

	switch usersChoice {
	case "1":
		if !utils.IsFileExist() {
			return
		}
		utils.FindAll()
	case "2":
		if !utils.IsFileExist() {
			return
		}
		phoneNumber = utils.ValidationLoop("Введите № телефона для поиска:", utils.PhoneFlag)
		utils.FindByNumber(phoneNumber)
	case "3":
		if _, err := os.Stat(utils.FileName); os.IsNotExist(err) {
			utils.CreateFile()
		}
		phoneNumber = utils.ValidationLoop("Введите № телефона:", utils.PhoneFlag)
		name = utils.ValidationLoop("Введите имя:", utils.NameFlag)
		if utils.AddContact(phoneNumber, name) {
			fmt.Println("Запись успешно добавлена")
		} else {
			fmt.Println("Запись не добавлена")
		}
	case "4":
		if !utils.IsFileExist() {
			return
		}
		phoneNumber = utils.ValidationLoop("Введите № телефона для изменения записи:", utils.PhoneFlag)
		if utils.EditContact(phoneNumber) {
			fmt.Println("Запись успешно изменена")
		} else {
			fmt.Println("Запись не изменена")
		}
	case "5":
		if !utils.IsFileExist() {
			return
		}
		phoneNumber = utils.ValidationLoop("Введите № телефона для удаления записи:", utils.PhoneFlag)
		if utils.DeleteContact(phoneNumber) {
			fmt.Println("Запись успешно удалена")
		} else {
			fmt.Println("Запись не удалена")
		}
	case "0":
		break
	default:
		fmt.Println("Неверный ввод. Введите цифру от 0 до 5")
	}
}

func printMenuAndSeparator() {
	fmt.Println("Выберите вариант действия:")
	fmt.Println("1 - показать все записи")
	fmt.Println("2 - поиск по номеру")
	fmt.Println("3 - добавить запись")
	fmt.Println("4 - изменить запись")
	fmt.Println("5 - удалить запись")
	fmt.Println("0 - выход")
	fmt.Println("----------------------------")
}
