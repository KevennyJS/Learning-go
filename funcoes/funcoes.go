package funcoes

import "fmt"

// func Nome (variavel Tipo) retorno
func Hello(name string) string {
	message := fmt.Sprintf("Olá, %v. Bem vindo!", name)
	return message
}
