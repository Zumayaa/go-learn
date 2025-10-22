package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("Escoge una opción: ")
	fmt.Scan(choice)
}

func edit() {
	var idGet int
	fmt.Print("Escoge un id para editar tarea: ")
	fmt.Scan(&idGet)

	for i, task := range tasks {
		if task.id == idGet {
			// Aquí va la edición
		}
	}
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
