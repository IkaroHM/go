package main

import "fmt"

func soma() {
	var n1, n2 float64 = 0, 0
	fmt.Println("Digite dois numeros para soma-los ")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Print("O resultado é ")
	fmt.Println(n1 + n2)
}
