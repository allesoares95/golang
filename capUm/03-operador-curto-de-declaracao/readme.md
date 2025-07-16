### Operador curto de declaração

- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas 
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
      
```go
	package main
	
	import "fmt"
	
	var y = "variável global"
	
	func main() {
	
	    x := 16
	
	    fmt.Printf("x: %v, %T\n", x, x)
	
	    fmt.Printf("y: %v, %T\n", y, y)
	
	    x, z := 20, "variável local"
	
	    fmt.Printf("x: %v, %T\n", x, x)
	
	    fmt.Printf("z: %v, %T\n", z, z)
	
	}
```
      
- Terminologia:
    - keywords (palavras-chave) são termos reservados
      [keywords]: 
		break        default      func         interface    select
			case         defer        go           map          struct
			chan         else         goto         package      switch
			const        fallthrough  if           range        type
			continue     for          import       return       var
      
    - operadores, operandos
		    binary_op  = "||" | "&&" | [rel_op](https://go.dev/ref/spec#rel_op) | [add_op](https://go.dev/ref/spec#add_op) | [mul_op](https://go.dev/ref/spec#mul_op) .
			rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
			add_op     = "+" | "-" | "|" | "^" .
			mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .
			unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
      
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
    - expressão -> qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
    - = para atribuir valores a variáveis já existentes
