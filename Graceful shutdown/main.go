package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	numChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(numChan)

		i := 1
		for {
			select {
			case numChan <- i * i:
				i++
			case <-exit:
				fmt.Println("Завершено")
				return
			}
		}
	}()

	for {
		select {
		case num, ok := <-numChan:
			if !ok {
				fmt.Println("Выход из программы")
				return
			}
			fmt.Println(num)
		case <-exit:
			close(numChan)
			wg.Wait()
			fmt.Println("Выход из программы")
			return
		}
	}
}
