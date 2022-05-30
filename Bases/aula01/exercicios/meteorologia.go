package exercicios

import "fmt"

func PrintMeteorologia() {
	fmt.Println("-------- Exercício 02 --------")
	var temperatura int16 = 14
	var umidade float32 = 0.66
	var pressaoAtmosferica float32 = 956.4

	fmt.Println("Temperatura: ", temperatura, " ˚C")
	fmt.Println("Umidade: ", umidade*100, "%")
	fmt.Println("Pressão atmosférica: ", pressaoAtmosferica, "hPa")
	fmt.Println()
}
