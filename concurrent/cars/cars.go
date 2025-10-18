package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	cars := []string{"Toyota Yarish", "Honda CRV", "Chevrolet Beat", "Ford Ranger", "Suzuki Swift"}

	var wg sync.WaitGroup
	var autolavados int = 2

	wg.Add(len(cars))

	for _, d := range cars {
		go func(name string) {
			defer wg.Done()

			for {

				mu.Lock()
				if autolavados > 0 {
					autolavados--
					mu.Unlock()

					fmt.Println(name, " lávandose")
					time.Sleep(2 * time.Second)
					fmt.Println(name, "se lavó")

					mu.Lock()
					autolavados++
					mu.Unlock()

					break

				} else {
					mu.Unlock()
					time.Sleep(10 * time.Millisecond)
				}
			}
		}(d)
	}

	wg.Wait()
	fmt.Println("Ya se lavaron todos los carros amor!!")
}
