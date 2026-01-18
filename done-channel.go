package main

import (
	"fmt"
	"time"
)

func doneChannel() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	done <- true
	fmt.Println("done")
}
