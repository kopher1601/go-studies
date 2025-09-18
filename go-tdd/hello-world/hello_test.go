package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		//  %q は値を二重引用符で囲む
		t.Errorf("got %q want %q", got, want)
	}
}
