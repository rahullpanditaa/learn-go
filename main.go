package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for range 12 {
		// i1 := <-chans[0]
		// i2 := <-chans[1]
		select {
		case m0 := <-chans[0]:
			fmt.Println("received", m0)
		case m1 := <-chans[1]:
			fmt.Println("received", m1)
		}
	}
}
