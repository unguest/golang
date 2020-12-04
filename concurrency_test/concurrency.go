package main

import "fmt"

func strlen(str string, length chan int) {
	length <- len(str)
}

func main() {
	str1 := "CONTAINS10" // Contains 10 chars
	str2 := "CONTAINS9" // Contains 9 chars
	length := make(chan int)

	go strlen(str2, length)
	go strlen(str1, length)

	/*
		When you use go, the functions seems to be executed in the reverse order...
	*/

	strlen1, strlen2 := <-length , <-length

	fmt.Println(strlen1)
	fmt.Println(strlen2)


	fmt.Print(strlen1 + strlen2)

}
