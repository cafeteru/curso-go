// Package say Modulo de prueba para comprobar la generación de docuimentación
package say

import "fmt"

// Greet manda un saludo a una persona
func Greet(name string) string {
	return fmt.Sprint("Welcome ", name)
}
