package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Docker!"
	got := Hello()

	if want != got {
		t.Errorf("Want:%s  Got:%s\n", want, got)
	}
}
