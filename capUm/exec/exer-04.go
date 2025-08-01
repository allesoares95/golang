// ### Na prática: exercício #4

// - Crie um tipo. O tipo subjacente deve ser int.
// - Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// - Na função main:
//     1. Demonstre o valor da variável "x"
//     2. Demonstre o tipo da variável "x"
//     3. Atribua 42 à variável "x" utilizando o operador "="
//     4. Demonstre o valor da variável "x"
// - Para os aventureiros: https://golang.org/ref/spec#Types
// - Agora já somos quase ninjas nível 1!
// - Solução: https://play.golang.org/p/snm4WuuYmG

package capUm

import "fmt"

type MyInt int

var testX MyInt

func ExerciseFour() {
	fmt.Printf("valor: %v\n", testX)
	fmt.Printf("tipo: %T\n", testX)

	testX = 42
	fmt.Printf("valor: %v\n", testX)
}
