package main

import "fmt"

func main() {

	var (
		choose int
	)

	directory := make(map[string]string)

	for choose != 5 {
		fmt.Println("")
		fmt.Println("Welcome to Zumayas Directory")
		fmt.Println("1- Consult the directory")
		fmt.Println("2- Add a new number")
		fmt.Println("3- Delete a number")
		fmt.Println("4- Modify a number")
		fmt.Println("5- Exit")
		fmt.Print("Make a choice: ")
		fmt.Scan(&choose)
		fmt.Println("")

		switch choose {
		case 1:
			consultNumber(directory)
		case 2:
			addNumber(directory)
		case 3:
			deleteNumber(directory)
		case 4:
			modifyNumber()
		}
	}
	fmt.Print("bye")
}

func consultNumber(directory map[string]string) {
	if len(directory) != 0 {
		for k, v := range directory {
			fmt.Println("Name:", k, "| Number:", v)
		}
	} else {
		fmt.Println("Empty xd")
	}
}

func addNumber(directory map[string]string) {
	var (
		name     string
		apellido string
	)

	fmt.Println("Write a name: ")
	fmt.Scan(&name)
	fmt.Println("Write a number: ")
	fmt.Scan(&apellido)
	directory[name] = apellido

}

func deleteNumber(directory map[string]string) {
	var (
		name string
	)

	if len(directory) != 0 {
		fmt.Println("Write a name to delete: ")
		fmt.Scan(&name)
		delete(directory, name)
	} else {
		fmt.Println("Empty xd")
	}
}

func modifyNumber() {
	fmt.Println("Here u can modify a number")
}
