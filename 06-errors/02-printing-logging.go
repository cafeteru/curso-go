package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var _, err = os.Open("No-existe.txt")
	if err != nil {
		fmt.Println("Ocurrio un error", err)
		log.Println("Ocurrio un error", err)
	}
}
