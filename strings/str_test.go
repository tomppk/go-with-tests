package str

import (
	"testing"
)

func TestLower(t *testing.T) {
	got := Lower("HELLO")
	expected := "hello"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func TestContain(t *testing.T) {

	t.Run("input contains substring", func(t *testing.T) {
		got := Contain("Hello, nice to meet you", "you")
		expected := true

		if got != expected {
			t.Errorf("got %t expected %t", got, expected)
		}

	t.Run("input does not contain substring", func(t *testing.T) {
		got := Contain("Hello, nice to meet you", "chicken")
		expected := false

		if got != expected {
			t.Errorf("got %t expected %t", got, expected)
		}
	})
	})
}

func TestJoined(t *testing.T) {
	s := []string{"over", "the", "rainbow"}
	got := Joined(s, " ")
	expected := "over the rainbow"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
