package main

import "testing"

func TestAddInt(t *testing.T) {
	testCase := []struct {
		Name     string
		Values   []int
		Expected int
	}{
		{"addInt() -> 0", []int{}, 0},
		{"addInt([]int{10,20,100}) -> 130", []int{10, 20, 100}, 130},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			sum := addInt(tc.Values...)
			if sum != tc.Expected {
				t.Error("%d != %d", sum, tc.Expected)
			} else {
				t.Log("%d == %d", sum, tc.Expected)
			}
		})
	}
}
