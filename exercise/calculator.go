package exercise

import "fmt"

// Calculadora de Pagos - Descuentos de Impuestos
// Precio % DA -> IV -> Total

func Calculator(price float64, discount float64, iva float64) (float64, error) {
	if price < 0 || discount < 0 || iva < 0 {
		return 0, fmt.Errorf("NEGATIVE VALUES")
	}

	var total float64 = price
	total -= total * discount
	total += total * iva

	return total, nil
}
