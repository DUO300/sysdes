package main

import "fmt"

/*
func out(n int) string {
	switch {
		case fifteen(n):
			return "FizzBuzz"
		case three(n):
			return "Fizz"
		case five(n):
			return "Buzz"
		default:
			return fmt.Sprintf("%d", n)
		}
}

func fifteen(n int) bool {
	return n % 15 == 0
}

func three(n int) bool {
	return n % 3 == 0
}

func five(n int) bool {
	return n % 5 == 0
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Printf("%s\n", out(i))
	}
}
*/

func IncrementV(v int) {
    v = v + 1
}
 
func IncrementP(p *int) {
    *p = *p + 1
}
 
func main() {
    v := 0
 
    IncrementV(v)
    fmt.Printf("v = %d\n", v)
 
    p := &v
    IncrementP(p)
    fmt.Printf("v = %d\n", v)
}
