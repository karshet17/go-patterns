package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Pipeline Pattern:")
	// pipelinePattern()

	fmt.Println("Done Channel:")
	//doneChannel()

	fmt.Println("Generator Pattern:")
	done := make(chan bool)
	genChan := repeatFunc(done)
	go func() {
		time.Sleep(time.Second * 2)
		close(done)
	}()
	for num := range genChan {
		fmt.Println(num)
	}

}
