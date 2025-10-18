package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	trabajadores := []string{"Trabajador 1", "Trabajador 2", "Trabajador 3", "Trabajador 4", "Trabajador 5"}

	wg.Add(len(trabajadores))
	var availableOvens int = 2
	var numBurger int = 3

	for _, d := range trabajadores {
		go func(name string) {
			defer wg.Done()

			for i := 1; i <= numBurger; i++ {
				mu.Lock() //bloquea hornos para primero verificar si hay disponibles y no pisotearse
				if availableOvens > 0 {
					availableOvens-- //si hay hornos, entonces resta 1
					mu.Unlock()      //como si hay entonces los desbloquea

					fmt.Println(name, " entra usando el horno, está cocinando su ", i, "burger")
					time.Sleep(1 * time.Second) // simula que cocina 1 hamburguesa
					fmt.Println(name, "sale del horno")

					mu.Lock() //se bloquea para uqe uno haga la suma primero porque si los dos saliesen al msimo tiempo se pisotearían
					availableOvens++
					mu.Unlock()
				} else {
					mu.Unlock() //si no hay hornos, tonces lo desbloquea para que si uno lo desocupa, otro lo agarre ne verguiza

					time.Sleep(10 * time.Millisecond) //espera pa volver a intentar
					i--                               //como no puedo hacer hamburguesa, la repite en la sig iteracion
				}
			}

		}(d)
	}
	wg.Wait() //el wait sirve para que el programa se termine solo si a huevo terminaron todos los hilos y pasaron por su wg.done
	fmt.Println("Todos terminaron de cocinar")

}
