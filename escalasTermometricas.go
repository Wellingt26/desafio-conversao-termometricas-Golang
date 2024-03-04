package main

import "fmt"

func main() {
	const pontoEbulicaoKelvin = 373.15

	pontoEbulicaoCelsius := pontoEbulicaoKelvin - 273.15

	fmt.Printf("O ponto de ebulição da água em Kelvin é: %g K, O ponto de ebulição da agua em Celsius é: %g °C\n", pontoEbulicaoKelvin, pontoEbulicaoCelsius)

}
