package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
	result := Hello("TOMOKI")
	want := "Hi, TOMOKI. Welcome!"
	if result != want {
		t.Errorf("fileWrite.Hello() = %q want %q", result, want)
	}
}
