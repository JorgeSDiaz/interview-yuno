package exercise

import (
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
		return false, fmt.Errorf("NOT A PALINDROME")
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
