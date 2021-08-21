package main

import (
	"fmt"
	"testing"
)

func TestBitVector(t *testing.T) {
	input := "game"
	if !BitVectorHasUniqueCharacters(input) {
		t.Error(fmt.Errorf("game should have unique characters"))
	}

	input = "food"
	if BitVectorHasUniqueCharacters(input) {
		t.Error(fmt.Errorf("food should not have unique characters"))
	}
}
