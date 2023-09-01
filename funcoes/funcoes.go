package funcoes

import (
	"errors"
	"fmt"
)

// func Nome (variavel Tipo) retorno
func Hello(name string) (string, error) {
	//Se o nome não for passado, ele retorna essa mensagem de erro
	if name == "" {
		return "", errors.New("Nome vazio")
	}

	//Se um nome for recebido, retornar o valor concatenado com a mensagem
	message := fmt.Sprintf("Olá, %v. Bem vindo!", name)
	return message, nil
}
