package main

import (
	"fmt"
)
func main() {
	fmt.Println("Ingrese la temperatura ambiental en Celsius: ")
	var temp float32
	fmt.Scanln(&temp)
		fmt.Printf("La temperatura en Farenheit es: ", temp*9/5 + 32.0)
}
