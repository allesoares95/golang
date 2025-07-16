// ### Na prática: exercício #3

// - Utilizando a solução do exercício anterior:
//     1. Em package-level scope, atribua os seguintes valores às variáveis:
//         1. para "x" atribua 42
//         2. para "y" atribua "James Bond"
//         3. para "z" atribua true
//     2. Na função main:
//         1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
//         2. Demonstre a variável "s".
// - Solução: https://play.golang.org/p/QFctSQB_h3

package capUm

import "fmt"

func ExerciseThree() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("x: %d, y: %s, z: %t", x, y, z)
	fmt.Println(s)
}

// Breakdown da linha: s := fmt.Sprintf("x: %d, y: %s, z: %t", x, y, z)
// 1. Operador de declaração curta (:=)
// Declara uma nova variável s e atribui um valor a ela
// O tipo da variável é inferido automaticamente (neste caso, string)
// Equivale a: var s string = fmt.Sprintf(...)

// 2. fmt.Sprintf
// Função do pacote fmt que formata uma string
// Similar ao printf em C, mas retorna a string formatada ao invés de imprimi-la
// Sprintf = "String Print Formatted"

// 3. String de formato com verbos
// "x: %d, y: %s, z: %t" é o template da string
// Os verbos de formatação (%d, %s, %t) são substituídos pelos valores das variáveis

// 4. Verbos de formatação:
// %d: formata um decimal (inteiro)
// %s: formata uma string
// %t: formata um boolean (true/false)

// 5. Argumentos
// x, y, z são as variáveis que serão inseridas nos respectivos verbos
// A ordem importa: primeiro argumento substitui o primeiro verbo, etc.

// Resultado:
// Se x = 42, y = "James Bond", z = true, então:

// Outros verbos comuns:
// %v: valor padrão (funciona com qualquer tipo)
// %f: números de ponto flutuante
// %x: hexadecimal
