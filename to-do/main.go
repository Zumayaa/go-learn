package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

var tasks []Task
var idList = 1

func main() {
	var choice int

	for {
		menu(&choice)

		switch choice {
		case 1:
			add()
		case 2:
			delete()
		case 3:
			edit()
		}
	}

}

func menu(choice *int) {
	if len(tasks) > 0 {
		fmt.Println("\nTareas ingresadas:")
		for _, task := range tasks {
			fmt.Printf("%d: %s (Completada: %t)\n", task.ID, task.Description, task.Done)
		}
	} else {
		fmt.Println("No hay tareas agregadas a la lista")
	}

	fmt.Println("--Biendosveni al To-Do list--")
	fmt.Println("1- Agregar tarea")
	fmt.Println("2- Eliminar tarea")
	fmt.Println("3- Editar tarea")
	fmt.Print("Escoge una opción: ")
	fmt.Scan(choice)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func edit() {
	var newDescription string
	scanner := bufio.NewScanner(os.Stdin)

	var idGet int
	fmt.Print("Escoge un id para editar tarea: ")
	scanner.Scan()
	idStr := scanner.Text()
	idGet, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	for i, task := range tasks {
		if task.ID == idGet {
			fmt.Print("Escribe una nueva descripción: ")
			scanner.Scan()
			newDescription = scanner.Text()
			if newDescription != "" {
				tasks[i].Description = newDescription
			} else {
				fmt.Println("No puede estar vacía")
			}
		}
	}
	fmt.Print("Se modificó la descripción de tu lista")
}

func delete() {
	fmt.Println("Wep")
}

func add() {
	var description string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Para salir de esta interfaz escribe 'quit': ")

	for {
		fmt.Println("Agrega una tarea negro")
		scanner.Scan()
		description = scanner.Text()

		if description == "quit" {
			break
		}

		if description != "" {
			task := Task{
				ID:          idList,
				Description: description,
				Done:        false,
			}
			tasks = append(tasks, task)
			idList++
		}
	}
}
