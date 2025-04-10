package hello

import "testing"

func TestHell(t *testing.T){
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}