package figura

type Cuadrado struct {
	alto  float64
	ancho float64
}

func (f *Cuadrado) Area() float64 {

	return f.ancho * f.alto

}
func (f *Cuadrado) Perimetro() float64 {

	return (2 * f.ancho) + (2 * f.alto)
}
func CrearCu(a, b float64) Cuadrado {
	fig := Cuadrado{}
	fig.alto = a
	fig.ancho = b
	return fig

}
