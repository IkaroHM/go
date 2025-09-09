package main

import "fmt"

func soma(a int, b int) int{
	return a+b
}

func main() {
	fmt.Println("total:", soma(10, 5))
	var array [3] string = [3]string{"teste", "teste", "teste"}
	fmt.Println(array)
	fmt.Println(cap(array))
	fmt.Println(len(array))
	/*
	var slice [2]string = [2]string{"teste", "teste", "teste"}
	*/

}