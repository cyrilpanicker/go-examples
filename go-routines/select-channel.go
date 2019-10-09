package main

import (
	"fmt"
	"time"
)

func sender(channel chan int, quit chan bool) {
	x := 0
	for {
		select {
		case channel <- x:
			x++
		case shouldQuit := <-quit:
			if shouldQuit {
				return
			}
		default:
			fmt.Println("no channel ready")
			time.Sleep(time.Millisecond)
		}
	}
}

func receiver(channel chan int, quit chan bool) {
	for index := 0; index < 5; index++ {
		fmt.Println(<-channel)
	}
	quit <- true
}

func selectChannel() {
	fmt.Println("-----------select channel starts-----------")
	channel := make(chan int)
	quit := make(chan bool)
	go sender(channel, quit)
	go receiver(channel, quit)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("-----------select channel ends-----------")
}
