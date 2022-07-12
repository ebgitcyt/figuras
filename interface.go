package figura

import "fmt"

type Figurita interface {
	Area() float64
	Perimetro() float64
}

func Calcular(objeto Figurita) {

	fmt.Println("El Area es:", objeto.Area())
	fmt.Println("El Perimetro es:", objeto.Perimetro())

}
