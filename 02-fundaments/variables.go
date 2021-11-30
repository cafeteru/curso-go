package main

import "fmt"

var (
	variableBoolean bool
	variableInt     int
	variableFloat64 float64
)

const constant = 5

func main() {
	fmt.Println(variableBoolean)
	variableBoolean = true
	fmt.Println(variableBoolean)
	variableInt = 1
	variableFloat64 = 3.1
	fmt.Println(variableInt, variableFloat64)

	s1 := "Hola mundo"
	s2 := `Esta es una línea 
	partida\n` // Respect all spaces, tabs and / or line breaks
	fmt.Printf(s1)
	fmt.Printf("El tipo de s1 es %T\n", s1)
	fmt.Printf(s2)
	fmt.Printf("El tipo de s2 es %T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("El tipo de bs es %T\n", bs)

	for i := 0; i < len(s1); i++ {
		valueUTF := s1[i]
		valueString := string(valueUTF)
		fmt.Println(i, valueString)
	}
	fmt.Println()
	for index, value := range s1 {
		valueString := string(value)
		fmt.Println(index, valueString)
	}
	fmt.Println()
	fmt.Println(constant)
}
