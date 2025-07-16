- Variável declarada em um code block é undefined em outro
- Para variáveis com uma abrangência maior, package level scope, utilizamos `var`
- Funciona em qualquer lugar
- Prestar atenção: chaves, colchetes, parênteses
  
```go
	package main

	import "fmt"
	
	func main() {
		y := 10
		qualquerCoisa(y)
	}
	
	func qualquerCoisa(x int) {
		fmt.Println(x)
	}

```
