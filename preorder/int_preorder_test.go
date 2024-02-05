package main

import (
	"reflect"
	"testing"
)

func TestGeneratePreorderRelation(t *testing.T) {
	testCases := []struct {
		name     string
		set      []int
		expected map[[2]int]bool
	}{
		{
			name: "empty set",
			set:  []int{},
			expected: map[[2]int]bool{},
		},
		{
			name: "single element",
			set:  []int{1},
			expected: map[[2]int]bool{
				{1, 1}: true,
			},
		},
		{
			name: "multiple elements",
			set:  []int{1, 2},
			expected: map[[2]int]bool{
				{1, 1}: true,
				{2, 2}: true,
				{1, 2}: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := generatePreorderRelation(tc.set)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("For set %v, expected %v, got %v", tc.set, tc.expected, actual)
			}
		})
	}
}
