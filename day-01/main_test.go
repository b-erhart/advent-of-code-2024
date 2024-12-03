package main

import "testing"

func TestTotalDistanceEmpty(t *testing.T) {
	left := []int{}
	right := []int{}

	dist := totalDistance(left, right)

	if dist != 0 {
		t.Fatalf("got %d, want 0", dist)
	}
}
