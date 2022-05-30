package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("/Users/sineto/Documents/bootcamp-dh/aula05/exercicios/customerss.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}

	conteudo := bufio.NewScanner(data)
	conteudo.Split(bufio.ScanLines)

	var linhas []string

	for conteudo.Scan() {
		linhas = append(linhas, conteudo.Text())
	}

	fmt.Println(linhas)
}
