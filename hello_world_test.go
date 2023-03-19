package task_0

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world"
	if got := HelloWorld(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
