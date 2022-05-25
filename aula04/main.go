package main

import (
	".bootcamp-dh/aula04/bases"
	"fmt"
	"log"
)

func main() {
	fmt.Println("-------- Aula 04 --------\n")
	fmt.Println("-------- Dia 01 --------\n")
	fmt.Println("-------- Exercicio 01 --------")

	msg, err := bases.VerificaSalario(14000)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(msg)
	}

	fmt.Println("\n-------- Exercicio 02 --------")
	msg2, err := bases.VerificaSalario2(16000)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(msg2)
	}

	fmt.Println("\n-------- Exercicio 03 --------")
	msg3, err := bases.VerificaSalario3(12000)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(msg3)
	}

	fmt.Println("\n-------- Exercicio 04 --------")
	salario, err := bases.CalculaSalario(75, 10)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(salario)
	}
}
