package stringutils

import "testing"

func TestReverse(t *testing.T) {
	result := Reverse("hello")
	expected := "olleh"
	if result != expected {
		t.Errorf("Reverse('hello') = %s; want %s", result, expected)
	}
}

func TestToUpperCase(t *testing.T) {
	result := ToUpperCase("hello")
	expected := "HELLO"
	if result != expected {
		t.Errorf("ToUpperCase('hello') = %s; want %s", result, expected)
	}
}
