- Conversão de tipos é o que soa.
- Em Go não se diz casting, se diz conversion.
- a = int(b)
- ref/spec#Conversions
- Fim da sessão. Parabéns! Dicas, motivação e exercícios.
  
```go
package main

import "fmt"

type hotdog int
var b hotdog = 101

func main() {
    x := 10
    fmt.Printf("%v\n", x)

    x = int(b)
    fmt.Printf("%v\n", x)
}
```
