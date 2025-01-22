package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Levi")
	want := "Hello, Levi!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
