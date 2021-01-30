package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины

	// вызываем горутины
	go work(1, &wg)
	go work(2, &wg)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшает внутренний счетчик активных элементов на единицу
	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}
