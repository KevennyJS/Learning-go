package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// func Nome (variavel Tipo) retorno
func Hello(name string) (string, error) {
	//Se o nome não for passado, ele retorna essa mensagem de erro
	if name == "" {
		return "", errors.New("Nome vazio")
	}

	//Se um nome for recebido, retornar o valor concatenado com a mensagem
	message := fmt.Sprintf(randomFormat(), name)
	//(Linha criada para rodar teste com falha)
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	//associando nomes com mensagens
	messages := make(map[string]string)

	//loop para o slice receber os nomes chamando a funcao hello e obtendo a mensagem de cada um
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		//asssociando mensagem recebida com o nome
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	//Criando o slice de format de mensagem
	formats := []string{
		"Olá, %v. Bem vindo!",
		"É ótimo te ver, %v!",
		"Salve, %v! Bom te encontrar!",
	}

	//Retornando de forma randomica um item da lista
	return formats[rand.Intn(len(formats))]
}
