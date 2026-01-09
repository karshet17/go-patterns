package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// stage 1: convert to channel
	numsChan := sliceToChannel(nums)

	// stage 2: double the numbers
	doubleChan := doubleNumbers(numsChan)

	for result := range doubleChan {
		fmt.Println(result)
	}

}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func doubleNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}
