package main

import "fmt"

type Entregador struct {
	nome string
	entregas int
	meta int
}

var entregador Entregador = Entregador{
	nome: "Ikaro",
	entregas: 5,
	meta: 4,
}


func (f Entregador) meta_entrega() {
	if f.entregas >= f.meta {
		fmt.Println(f.nome, "Bateu a meta!")
	} else {
		fmt.Println(f.nome, "Nao bateu a meta!")
	}
}

func main() {
	entregador.meta_entrega()
}