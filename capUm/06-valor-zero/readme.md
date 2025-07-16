- Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
- O que é valor zero?
- Os zeros:
    - ints: 0
    - floats: 0.0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível.
- Use var para package-level scope.

```go
	package main
	
	import "fmt"
	
	var a int
	var b float64
	var c string
	var d bool  
	
	func main() {
	    fmt.Printf("%v, %T\n", a, a)
	    fmt.Printf("%v, %T\n", b, b)
	    fmt.Printf("%v, %T\n", c, c)
	    fmt.Printf("%v, %T\n", d, d)
	}
```

```cli
0, int
0, float64
, string
false, bool
```
