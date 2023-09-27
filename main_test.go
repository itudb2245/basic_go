package main

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {
	actual := main.Sum(1, 2)
	expected := 3
	if actual != expected {
		t.Errorf("Sum(1, 2): actual %v, expected %v", actual, expected)
	}
}

func Sum(a int, b int) int {
	return a + a
}