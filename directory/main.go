package main

import "fmt"

func main() {

	var (
		choose int
	)

	directory := make(map[string]string)

	for choose != 5 {
		fmt.Println("")
		fmt.Println("Welcome to Zumayas Directory \n 1- Consult the directory \n 2- Add a new number \n 3- Delete a number \n 4- Modify a number \n 5- Exit")
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
			modifyNumber(directory)
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

func modifyNumber(directory map[string]string) {
	var (
		name, newNumber, newName, oldName string
		choice                            int
	)

	if len(directory) == 0 {
		fmt.Println("Empty")
	} else {
		fmt.Print("What do you want to modify: \n 1- Name \n 2- Number \n Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 2:
			fmt.Print("Type the name: ")
			fmt.Scan(&name)
			if _, exists := directory[name]; exists { //esta siempre es la sintaxis para saber si algo es true o false, asi creas un bool pa preguntar
				fmt.Print("Write the new number: ")
				fmt.Scan(&newNumber)
				directory[name] = newNumber //solamente recorres el nombre con el bool (es falso de serie), si lo encuentra (se hace true)
				// sigue el codigo, con la variable nueva le asignas a la key el nuevo numero
				fmt.Println("Number updated!")
			} else {
				fmt.Println("Name not found.")
			}
		case 1:
			fmt.Print("Type the name: ")
			fmt.Scan(&oldName)
			//aqui le dices a go que oldname es la key y lo que quieres obtener es el value, la segunda posicion, busca el name y dame el value
			if value, exists := directory[oldName]; exists { //esta siempre es la sintaxis para saber si algo es true o false, asi creas un bool pa preguntar
				fmt.Print("Write the new name: ")
				fmt.Scan(&newName)
				directory[newName] = value
				delete(directory, oldName)
				fmt.Println("Name updated!")
			} else {
				fmt.Println("Name not found.")
			}
		default:
			fmt.Println("Write a real thing tontopollas")
		}
	}

}
