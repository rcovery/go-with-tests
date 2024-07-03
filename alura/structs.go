package main

import (
	"fmt"

	r "tests.com/restaurantes"
)

func structs() {
	comida1 := r.MarmitaNegaMaluca{
		Comida:     "Parmegiana",
		Valor:      15,
		Entregador: true,
	}

	fmt.Println(comida1)

	comida1.AlterarPreco(20)
	Entrega(&comida1)

	fmt.Println(comida1)

	// comida2 := MarmitaNegaMaluca{
	// 	"Panqueca", 18, true,
	// }

	// var comida3 *MarmitaNegaMaluca
	// comida3 = new(MarmitaNegaMaluca)
	// comida3.Comida = "Feijoada"
	// comida3.Valor = 20
}

func Entrega(comida r.TemEntregador) {
	comida.Entregar()
}
