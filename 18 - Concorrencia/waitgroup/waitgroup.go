package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Concorrencia != paralelismo

	var waitGroup sync.WaitGroup

	waitGroup.Add(3) //3 goroutines que fazem parte deste grupo de espera

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("GoRoutine 3")
		waitGroup.Done() //-1
	}()
	waitGroup.Wait() //espera as contagens das goroutines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
