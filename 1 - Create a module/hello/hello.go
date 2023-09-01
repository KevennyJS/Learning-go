package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	//Definindo propriedades do Logger
	log.SetPrefix("Saudações:")
	log.SetFlags(0)

	//Solicitando mensagem de saudação
	message, err := greetings.Hello("Kevenny")

	//se retornar um erro, printe no console e feche o programa
	if err != nil {
		log.Fatal(err)
	}

	//se não retornar erro, printe a mensagem no connsole
	fmt.Println(message)

	//========================================================

	//slice de nomes
	nomes := []string{"Kevenny", "Joanne", "Ericles", "Milena", "Kendy"}

	//solicitando mensagem de saudação para os nomes
	messages, err := greetings.Hellos(nomes)
	if err != nil {
		log.Fatal(err)
	}

	//se nenhum erro for encontrado, printe as mensagens
	fmt.Println(messages)
}
