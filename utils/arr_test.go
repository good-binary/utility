package utils

import (
	"testing"
)

func TestSlicer(t *testing.T) {
	// Initialize a new Slicer
	s := NewSlicer([]int{1, 2, 3})

	// Test Len method
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test Get method
	if s.Get(1) != 2 {
		t.Errorf("Expected 2 at index 1, got %d", s.Get(1))
	}

	// Test Set method
	s.Set(1, 4)
	if s.Get(1) != 4 {
		t.Errorf("Expected 4 at index 1, got %d", s.Get(1))
	}

	// Test Append method
	s.Append(5)
	if s.Get(3) != 5 {
		t.Errorf("Expected 5 at index 3, got %d", s.Get(3))
	}

	// Test Prepend method
	s.Prepend(0)
	if s.Get(0) != 0 {
		t.Errorf("Expected 0 at index 0, got %d", s.Get(0))
	}

	// Test Remove method
	s.Remove(0)
	if s.Get(0) != 1 {
		t.Errorf("Expected 1 at index 0, got %d", s.Get(0))
	}

	// Test Clear method
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}
}
