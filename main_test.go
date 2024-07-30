package main

import (
	"testing"
)

// Use Table Driven Design. Encontré que era una manera fácil de ejecutar varios test unicamente definiendo el arreglo de casos a probar y me gusto
// Recursos que use:
// 					- https://medium.com/@rasheed99/how-to-write-table-driven-tests-in-go-8e96ef048cca
// 					- https://pkg.go.dev/testing
// 					- https://blog.friendsofgo.tech/posts/empezando-con-los-tests-automatizados-en-go/

// Deje los test de ambos en el mismo archivo por hacerlo rápido, había pensado en dejar un paquete de "exercise" las funciones y en otro paquete "test"
// los test para cada uno de las funciones.
// ├── excercise
// │   ├── calculator.go
// │   └── palindrome.go
// ├── go.mod
// ├── main.go
// └── test
//    └── main_test.go

// Ejecutar con:
// 				- go test -v -run TestIsPalindrome
// 				- go test -v -run TestCalculator
// Para ver el coverage -> go test -cover

// Test Palindrome
type IsPalindromeTestCase struct {
	name          string
	word          string
	expectedValue bool
	expectedError bool
}

var isPalindromeTests = []IsPalindromeTestCase{
	{"'hello world' is palindrome ?", "hello world", false, false},
	{"'reconocer' is palindrome ?", "reconocer", true, false},
	{"'Reconocer' is palindrome ?", "Reconocer", true, false},
	{"'' is palindrome ?", "", false, true},
	{"'a' is palindrome ?", "a", true, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, testCase := range isPalindromeTests {
		t.Run(testCase.name, func(t *testing.T) {
			result, error := IsPalindrome(testCase.word)
			// Si hay error que no se espera
			if (error != nil) != testCase.expectedError {
				t.Fatalf("IsPalindrome(%v): Unexpected Error -> %v", testCase.word, error)
			}
			// Si no hay error pero el resultado no es el esperado
			if error == nil && result != testCase.expectedValue {
				t.Fatalf("IsPalindrome(%v): Unexpected Return -> %t", testCase.word, result)
			}
			// Si hay error y no es el esperado
			if error != nil && error.Error() != "NOT A WORD" {
				t.Fatalf("IsPalindrome(%v): Unexpected Error Message -> %v", testCase.word, error)
			}
		})
	}
}

// Calculator Test
type CalculatorTestCase struct {
	name                 string
	price, discount, iva float64
	expectedValue        float64
	expectedError        bool
}

var CalculatorTests = []CalculatorTestCase{
	{"Calculator(100.0, 0.1, 0.01)", 100.0, 0.1, 0.01, 90.9, false},
	{"Calculator(100.0, 0.0, 0.01)", 100.0, 0.0, 0.01, 101.0, false},
	{"Calculator(100.0, -0.1, 0.01)", 100.0, -0.1, 0.01, 0, true},
	{"Calculator(-100.0, 0.1, 0.01)", -100.0, 0.1, 0.01, 0, true},
	{"Calculator(100.0, 0.1, -0.01)", 100.0, 0.1, -0.01, 0, true},
}

func TestCalculator(t *testing.T) {
	for _, testCase := range CalculatorTests {
		t.Run(testCase.name, func(t *testing.T) {
			result, error := Calculator(testCase.price, testCase.discount, testCase.iva)
			if (error != nil) != testCase.expectedError {
				t.Fatalf("%v: Unexpected Error -> %v", testCase.name, error)
			}
			if error == nil && result != testCase.expectedValue {
				t.Fatalf("%v: Unexpected Return -> %v", testCase.name, result)
			}
			if error != nil && error.Error() != "NEGATIVE VALUES" {
				t.Fatalf("%v: Unexpected Error Message -> %v", testCase.name, error)
			}
		})
	}
}
