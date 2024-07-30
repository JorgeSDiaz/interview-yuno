package main

import (
	"errors"
	"fmt"
	"strings"
)

// Objetivo: Detectar palíndromo
// Entrada: "hola mundo" - "Reconocer"
// Salida: false - true

// Case sensitive : false
// Phrases : replace " " -> ""

func IsPalindrome(word string) (bool, error) {
	if word == "" {
		return false, errors.New("NOT A WORD")
	}

	word = strings.ToLower(strings.ReplaceAll(word, " ", ""))

	var word_letters []rune = []rune(word)
	var word_letters_reverse []rune

	for i := 0; i < len(word_letters); i++ {
		word_letters_reverse = append(word_letters_reverse, word_letters[len(word_letters)-i-1])
	}

	if word != string(word_letters_reverse) {
		return false, nil
	}

	return true, nil
}

// Calculadora de Pagos - Descuentos de Impuestos
// Precio % DA -> IV -> Total

func Calculator(price float64, discount float64, iva float64) (float64, error) {
	if price < 0 || discount < 0 || iva < 0 {
		return 0, errors.New("NEGATIVE VALUES")
	}

	var total float64 = price
	total -= total * discount
	total += total * iva

	return total, nil
}

func main() {
	fmt.Println(IsPalindrome("Reconocer"))
	fmt.Println(IsPalindrome("hola mundo"))
	fmt.Println(Calculator(100.0, 0, 0.01))
	fmt.Println(Calculator(-100.0, 0.1, 0.01))
}
