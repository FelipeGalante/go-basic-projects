package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello world, from Gopher."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
