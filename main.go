package main

import (
	"fmt"

	"github.com/Meryyy2025/godesde0/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" || os == "OS. X" {
		fmt.Println("Esto no es Windows, es", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
	num, text := ejercicios.ValidateTextNumberValue("fff")
	fmt.Println(num)
	fmt.Println(text)
}
