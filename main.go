package main

import (
	"fmt"
	"strings"
)

// Go calculator base on Prefix Evaluation

func main() {
	fmt.Println("Prefix calculator")
	fmt.Println("Tokens are separated by whitespaces.")
	fmt.Println("Operators:")
	fmt.Println("  Binary: + - * / ^")
	var task *Conversion = getConversion()
	fmt.Print("> ")
	input := ""
	fmt.Scanln(&input)
	prefix_result := task.infixToPrefix(input)
	sum := 0.0
	sum, _ = Prefix(strings.NewReader(prefix_result))
	fmt.Printf(" sum is : %.16g\n", sum)

	// Test
	//(2+3)
	//result :  5
	//((1*3)+(4-5))
	//result : -2.333333333333333
	//((1/(2-3+4))*(5-6)*7)
	//result : 2
}
