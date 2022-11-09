package main

import (
	"reflect"
	"testing"
)

func TestLinearTableAppendOne1(t *testing.T) {
	original := []int{90, 70, 50, 80, 60, 85}
	appendValue := 75
	expected := []int{90, 70, 50, 80, 60, 85, 75}

	result := linearTableAppendOne(original, appendValue)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}

func TestLinearTableAppendOne2(t *testing.T) {
	original := []int{90, 70, 50, 80, 60, 85}
	appendValue := 75
	expected := []int{90, 70, 50, 80, 60, 85, 75}

	result := linearTableAppendOne(original, appendValue)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}
