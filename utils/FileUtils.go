package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const FileName string = "phoneNumbers.txt"
const PhoneFlag string = "phone"
const NameFlag string = "name"

var PhoneRegexp *regexp.Regexp = regexp.MustCompile("^[+]380([0-9]{9})$")
var NameRegexp *regexp.Regexp = regexp.MustCompile("^[0-9A-Za-z]{3,15}$")

func CreateFile() {
	f, err := os.Create(FileName)
	check(err)
	defer f.Close()
}

func AddContact(number, name string) bool {
	f, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	f.WriteString(number + ":" + name + "\n")
	return true
}

func FindAll() {
	f, err := ioutil.ReadFile(FileName)
	check(err)
	if isFileEmpty(len(f)) {
		return
	}
	fmt.Println("Найдены записи:")
	fmt.Println(string(f))
}

func FindByNumber(number string) {
	f, err := ioutil.ReadFile(FileName)
	check(err)
	rows := strings.Split(string(f), "\n")

	if isFileEmpty(len(rows)) {
		return
	}

	for _, row := range rows {
		if strings.Contains(row, number) {
			fmt.Println("Найдена запись:", row)
			return
		}
	}
	fmt.Println("Запись с таким номером не существует")
}

func EditContact(number string) bool {
	f, err := ioutil.ReadFile(FileName)
	check(err)
	rows := strings.Split(string(f), "\n")

	if isFileEmpty(len(rows)) {
		return false
	}

	var name string
	var isDataExist bool
	var rowItem []string

	for idx, row := range rows {
		rowItem = strings.Split(row, ":")
		if strings.EqualFold(rowItem[0], number) {
			isDataExist = true
			fmt.Println("Найдена запись:", row)
			name = ValidationLoop("Введите новое имя:", NameFlag)
			rowItem[1] = name
			rows[idx] = strings.Join(rowItem, ":")
			output := strings.Join(rows, "\n")
			err = ioutil.WriteFile(FileName, []byte(output), 0644)
			check(err)
			return true
		}
	}

	if !isDataExist {
		fmt.Println("Запись с таким номером не существует")
	}
	return false
}

func DeleteContact(number string) bool {
	f, err := ioutil.ReadFile(FileName)
	check(err)
	rows := strings.Split(string(f), "\n")

	if isFileEmpty(len(rows)) {
		return false
	}

	var userChoice string
	var isDataExist bool
	var rowItem []string
	var changedRows []string

	for idx, row := range rows {
		rowItem = strings.Split(row, ":")
		if strings.EqualFold(rowItem[0], number) {
			isDataExist = true
			fmt.Println("Найдена запись:", row)
			for {
				fmt.Println("Удалить запись? Для подтверждения нажмите 1. Для отмены - 0")
				fmt.Scanln(&userChoice)
				switch userChoice {
				case "0":
					return false
				case "1":
					changedRows = append(changedRows, rows[:idx]...)
					idx++
					changedRows = append(changedRows, rows[idx:]...)
					output := strings.Join(changedRows, "\n")
					err = ioutil.WriteFile(FileName, []byte(output), 0644)
					check(err)
					return true
				default:
					fmt.Println("Неверный ввод. Введите цифру 1 или 0")
				}
			}
		}
	}

	if !isDataExist {
		fmt.Println("Запись с таким номером не существует")
	}
	return false
}

func isFileEmpty(length int) bool {
	// Используется 2, т.к. в файле есть пустая строка (перенос)
	if length < 2 {
		fmt.Println("В файле нет записей")
		return true
	}
	return false
}

func IsFileExist() bool {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		fmt.Println("Файл с номерами пока отсутствует. Для его создания совершите добавление записи")
		return false
	}
	return true
}

func validate(data string, flag string) bool {
	if flag == PhoneFlag {
		return PhoneRegexp.MatchString(data)
	}
	if flag == NameFlag {
		return NameRegexp.MatchString(data)
	}
	return false
}

func ValidationLoop(msg string, flag string) string {
	var data string
	for {
		fmt.Println(msg)
		fmt.Scanln(&data)
		if validate(data, flag) {
			break
		}
		fmt.Println("Введите верные данные")
		fmt.Println("Формат номера телефона +380xxYYYYYYY")
		fmt.Println("Формат имени - только буквы A-z и/или цифры (от 3 до 15 символов)")
	}
	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
