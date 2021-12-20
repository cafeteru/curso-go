// Package say Modulo de prueba para comprobar la generación de docuimentación
package say

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	var result = "Welcome Iván"
	var s = Greet("Iván")
	if s != result {
		t.Error("Expected", result, "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Iván"))
	//Output:
	//"Welcome Iván"
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Iván")
	}
}
