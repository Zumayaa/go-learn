package main

import "fmt"

func main() {
	fmt.Println("Wep")

	var (
		//name   string
		//number string
		choose int
	)

	directory := make(map[string]string)

	for choose != 5 {
		fmt.Println("Welcome to Zumayas Directory")
		fmt.Println("1- Consult the directory")
		fmt.Println("2- Add a new number")
		fmt.Println("3- Delete a number")
		fmt.Println("4- Modify a number")
		fmt.Println("5- Exit")
		fmt.Print("Make a choose: ")
		fmt.Scan(&choose)

		switch choose {
		case 1:
			consultNumber(directory)
		case 2:
			addNumber()
		case 3:
			deleteNumber()
		case 4:
			modifyNumber()
		}
	}
	fmt.Print("bye")
}

func consultNumber(directory map[string]string) {
	for k, v := range directory {
		fmt.Println("Nombre:", k, "| Número:", v)
	}
}

func addNumber() {
	fmt.Println("Here u can add a new number")
}

func deleteNumber() {
	fmt.Println("Here u can delete a number")
}

func modifyNumber() {
	fmt.Println("Here u can modify a number")
}
