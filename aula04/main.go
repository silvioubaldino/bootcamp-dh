package main

import (
	".bootcamp-dh/aula04/bases"
	"fmt"
	"log"
)

func main() {
	fmt.Println("-------- Aula 04 --------\n")
	fmt.Println("-------- Dia 01 --------\n")
	fmt.Println("-------- Exercicio 01 --------\n")

	msg, err := bases.VerificaSalario(14000)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg)

	fmt.Println("-------- Exercicio 02 --------\n")
	msg2, err := bases.VerificaSalario2(16000)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg2, "\n")

	fmt.Println("-------- Exercicio 03 --------\n")
	msg3, err := bases.VerificaSalario3(12000)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(msg3)
}
