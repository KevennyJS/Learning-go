package main

import (
	"example.com/funcoes"
	"fmt"
	"log"
)

func main() {
	//Definindo propriedades do Logger
	log.SetPrefix("Saudações:")
	log.SetFlags(0)

	//Solicitando mensagem de saudação
	message, err := funcoes.Hello("Kevenny")

	//se retornar um erro, printe no console e feche o programa
	if err != nil {
		log.Fatal(err)
	}

	//se não retornar erro, printe a mensagem no connsole
	fmt.Println(message)
}
