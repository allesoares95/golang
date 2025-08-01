// ### Na prática: exercício #5

// - Utilizando a solução do exercício anterior:
//     1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//     2. Na função main:
//         1. Isto já deve estar feito:
//             1. Demonstre o valor da variável "x"
//             2. Demonstre o tipo da variável "x"
//             3. Atribua 42 à variável "x" utilizando o operador "="
//             4. Demonstre o valor da variável "x"
//         2. Agora faça tambem:
//             1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//             2. Demonstre o valor de "y"
//             3. Demonstre o tipo de "y"
// - Solução: https://play.golang.org/p/uq81T_fasP

package capUm

import "fmt"

type Inteiro int

var testeX Inteiro
var testeY int

func ExerciseFive() {
	fmt.Println("Valor de testeX:", testeX)
	fmt.Printf("Tipo de testeX: %T\n", testeX)

	testeX = 42
	fmt.Println("Valor de testeX após atribuição:", testeX)

	testeY = int(testeX)
	fmt.Println("Valor de testeY:", testeY)
	fmt.Printf("Tipo de testeY: %T\n", testeY)
}
