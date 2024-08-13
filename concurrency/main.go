package main

import "fmt"

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	// call goroutines to fill the channels with ints
	go nums1(channel1)
	go nums2(channel2)

	// read from the channels, will display in order
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
}
