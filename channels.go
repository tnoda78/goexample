package main

import "fmt"

func main() {

	mesages := make(chan string)

	go func() { mesages <- "ping" }()

	msg := <-mesages
	fmt.Println(msg)
}
