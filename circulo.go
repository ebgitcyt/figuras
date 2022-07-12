package figura

const pi = 3.141516

type Circulo struct {
	radio float64
}

func (f *Circulo) Area() float64 {

	return pi * (f.radio * f.radio)

}
func (f *Circulo) Perimetro() float64 {

	return 2 * pi * f.radio
}
func CrearC(a float64) Circulo {
	fig := Circulo{}
	fig.radio = a
	return fig
}
