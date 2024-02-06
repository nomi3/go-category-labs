package main

import (
	"reflect"
	"testing"
)

func TestgenerateAnyPreorderRelationInts(t *testing.T) {
	set := []interface{}{1, 2, 3}
	expected := map[[2]interface{}]bool{
		{1, 1}: true, {1, 2}: true, {1, 3}: true,
		{2, 2}: true, {2, 3}: true,
		{3, 3}: true,
	}
	preorder := generateAnyPreorderRelation(set)

	if !reflect.DeepEqual(preorder, expected) {
		t.Errorf("Expected preorder relation did not match. Got %v, want %v", preorder, expected)
	}
}

func TestgenerateAnyPreorderRelationStrings(t *testing.T) {
	set := []interface{}{"apple", "banana", "cherry"}
	expected := map[[2]interface{}]bool{
		{"apple", "apple"}: true, {"apple", "banana"}: true, {"apple", "cherry"}: true,
		{"banana", "banana"}: true, {"banana", "cherry"}: true,
		{"cherry", "cherry"}: true,
	}
	preorder := generateAnyPreorderRelation(set)

	if !reflect.DeepEqual(preorder, expected) {
		t.Errorf("Expected preorder relation did not match. Got %v, want %v", preorder, expected)
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		x        interface{}
		y        interface{}
		expected bool
	}{
		{1, 2, true},
		{2, 1, false},
		{"apple", "banana", true},
		{"banana", "apple", false},
		{1, "apple", false},
	}

	for _, test := range tests {
		result := compare(test.x, test.y)
		if result != test.expected {
			t.Errorf("compare(%v, %v) = %t; want %t", test.x, test.y, result, test.expected)
		}
	}
}
