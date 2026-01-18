package main

import "math/rand"

func repeatFunc(done <-chan bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(100):
			}
		}
	}()

	return out
}
