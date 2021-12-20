package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var answer string

func main() {
	example1()
	example2()
	example3()
}

func example1() {
	var value, err = fmt.Println("Hola")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func example2() {
	fmt.Println("Nombre: ")
	var _, err = fmt.Scan(&answer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Su nombre es", answer)
}

func example3() {
	var file, err = os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer -> pospone su ejecución en una pila que no se ejecuta hasta termine main
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	newContent := strings.NewReader(answer)
	_, err = io.Copy(file, newContent)
	if err != nil {
		return
	}
}
