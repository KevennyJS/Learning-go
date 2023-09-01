package greetings

import (
	"regexp"
	"testing"
)

// Testando chamada de greetings.Hello com nome, verificando validade do retorno
func TestHelloName(t *testing.T) {
	name := "João"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("João")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("João") %q, %v, quero encontrar por %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty chama grettings.Hello com um valor vazio e verifica se ocorre um erro
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Olá ("") =%q, %v, wait "", erro`, msg, err)
	}
}
