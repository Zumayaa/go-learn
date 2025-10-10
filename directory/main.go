package main

import "fmt"

func main() {
	fmt.Println("Wep")

	var (
		//name   string
		//number string
		choose int
	)

	//directory := make(map[string]string)

	for choose != 4 {
		fmt.Println("Welcome to Zumayas Directory")
		fmt.Println("1- Add a new number")
		fmt.Println("2- Delete a number")
		fmt.Println("3- Modify a number")
		fmt.Println("4- Exit")
		fmt.Print("Make a choose: ")
		fmt.Scan(&choose)

		switch choose {
		case 1:
			addNumber()
		case 2:
			deleteNumber()
		case 3:
			modifyNumber()
		}
	}
	fmt.Print("bye")
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
