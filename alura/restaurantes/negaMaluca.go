package restaurantes

import "fmt"

type MarmitaNegaMaluca struct {
	Comida     string
	Valor      float32
	Entregador bool
}

func (m *MarmitaNegaMaluca) AlterarPreco(newValue float32) {
	fmt.Println("Alterando preço da", m.Comida, "de", m.Valor, "para", newValue)
	m.Valor = newValue
}

func (m *MarmitaNegaMaluca) Entregar() {
	fmt.Println("Entregando", m.Comida)
}
