package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	if s := Greet("James"); s != "Welcome James" {
		t.Error("got", s, "Expected", "Welcome James")
	}

}

func ExampleGreet() {
	fmt.Println(Greet("James"))
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
