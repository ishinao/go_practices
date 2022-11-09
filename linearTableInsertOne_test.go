package main

import (
	"reflect"
	"testing"
)

func TestLinearTableInsertOne1(t *testing.T) {
	original := []int{90, 70, 50, 80, 60, 85}
	insertValue := 75
	insertPosition := 2
	expected := []int{90, 70, 75, 50, 80, 60, 85}
	result := linearTableInsertOne(original, insertValue, insertPosition)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}
func TestLinearTableInsertOne2(t *testing.T) {
	original := []int{90, 70, 60, 85}
	insertValue := 23
	insertPosition := 3
	expected := []int{90, 70, 60, 23, 85}
	result := linearTableInsertOne(original, insertValue, insertPosition)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}
