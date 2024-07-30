package main

import (
	"fmt"

	"github.com/JorgeSDiaz/interview-yuno/excersice/calculator"
	"github.com/JorgeSDiaz/interview-yuno/excersice/palindrome"
)

func main() {
	fmt.Println(palindrome.IsPalindrome("Reconocer"))
	fmt.Println(palindrome.IsPalindrome("hola mundo"))
	fmt.Println(calculator.Calculator(100.0, 0.1, 0.01))
	fmt.Println(calculator.Calculator(-100.0, 0.1, 0.01))
}
