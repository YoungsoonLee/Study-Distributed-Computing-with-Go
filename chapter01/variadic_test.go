package main

import "testing"

func TestSimpleVariadicToSlice(t *testing.T) {
	// test for no arguments
	if val := simpleVariadicToSlice(); val != nil {
		t.Error("value should be nil", nil)
	} else {
		t.Log("simpleVariadicToSlice() -> nil")
	}

	// test for random set of values
	vals := simpleVariadicToSlice(1, 2, 3)
	expected := []int{1, 2, 3}
	isErr := false
	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}

	if isErr {
		t.Error("value should be []int{1,2,3}", vals)
	} else {
		t.Log("simpleVariadicToSlice(1,2,3) -> []int{1,2,3}")
	}

	// test for a slice
	vals = simpleVariadicToSlice(expected...)
	isErr = false
	for i := 0; i < 3; i++ {
		if vals[i] != expected[i] {
			isErr = true
			break
		}
	}

	if isErr {
		t.Error("value should be []int{1,2,3}", vals)
	} else {
		t.Log("simpleVariadicToSlice([]int{1,2,3}...) -> []int{1,2,3}")
	}
}

func TestMixedVariadicToSlice(t *testing.T) {
	// test for simple argument & no variadic arguments
	name, numbers := mixedVariadicToSlice("Bob")
	if name == "Bob" && numbers == nil {
		t.Log("received as expected: Bob, <nil slice>")
	} else {
		t.Error("received unexpected values: %s %s", name, numbers)
	}
}
