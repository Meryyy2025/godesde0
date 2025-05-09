package ejercicios

import (
	"strconv"
)

// This function receive param textNumber, parse this to int value and return a
// specific message if this value is mayor/minor than onehundred
func ValidateTextNumberValue(textNumber string) (int, string) {
	// using ParseInt method

	entero, error := strconv.Atoi(textNumber)
	if error != nil {
		return 0, "Hubo un error" + error.Error()
	}
	if entero > 100 {
		return entero, "Es mayor a 100"
	} else {
		return entero, "Es menor a 100"
	}

}
