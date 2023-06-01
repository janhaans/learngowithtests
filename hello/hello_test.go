package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jan")
	want := "Hello Jan!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
