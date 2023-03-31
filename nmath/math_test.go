package nmath

import "testing"

func TestAdd(t *testing.T) {
	t.Log("TestAdd")

	result := Add(10, 10)
	expected := 20

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	t.Log("TestSubtract")
	result := Subtract(10, 10)
	expected := 0

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	t.Log("TestMultiply")
	result := Multiply(10, 10)
	expected := 100

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	t.Log("TestMultiply")
	result := Divide(10, 10)
	expected := 1

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}
