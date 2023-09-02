package reverse

import "strconv"

// retorna o recimal reverso de um inteiro i
func Int(i int) int {
	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
	return i
}
