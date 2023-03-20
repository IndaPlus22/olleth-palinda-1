package main

import (
	"testing"
)

// test that ConcurrentSum sums an even-length array correctly
func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test that ConcurrentSum sums a single-element array correctly
func TestSumConcurrentCorrectlySumsSingleElementArray(t *testing.T) {
	arr := []int{42}
	expected := 42

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}

// test that ConcurrentSum sums a large array correctly
func TestSumConcurrentCorrectlySumsLargeArray(t *testing.T) {
	arr := make([]int, 1000000)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	expected := 500000500000

	actual := ConcurrentSum(arr)

	if actual != expected {
		t.Errorf("expected %d, was %d", expected, actual)
	}
}
