package basics

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	if result != 10 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
