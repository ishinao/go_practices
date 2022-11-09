package main

import (
	"reflect"
	"testing"
)

func TestLinearTableRemoveOne1(t *testing.T) {
	original := []int{90, 70, 50, 80, 60, 85}
	removePosition := 2
	expected := []int{90, 70, 80, 60, 85}
	result := linearTableRemoveOne(original, removePosition)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}
