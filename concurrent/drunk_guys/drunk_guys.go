package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex // la puerta del baño

	// Creamos dos borrachos
	borrachos := []string{"Borrachín 1", "Borrachín 2"}

	var wg sync.WaitGroup
	wg.Add(len(borrachos)) //agrega dos borrachos

	for _, b := range borrachos {
		go func(name string) {
			defer wg.Done() //defer es como el default del switch, esto se pone para ejecutarse al final sin importar que pase, es como si un hilo dijera "acabé!"

			// Qquieren entrar al baño
			mu.Lock() //la puerta del baño y con esto das a entender que solo uno puede entrar para que netre hilos no se pise
			fmt.Println(name, "entra al baño ")
			time.Sleep(2 * time.Second) //simulacion
			fmt.Println(name, "sale del baño ")
			mu.Unlock() //desbloqueas la puerta para que otro hilo pase

		}(b) //se itera la variable sin que se cree de nuevo
	}

	wg.Wait() //el wait sirve para que el programa se termine solo si a huevo terminaron todos los hilos y pasaron por su wg.done
	fmt.Println("Todos terminaron de usar el baño")
}
