package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort1(t *testing.T) {
	original := []int{60, 50, 95, 80, 70}
	expected := []int{50, 60, 70, 80, 95}
	result := bubbleSort(original)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}

func TestBubbleSort2(t *testing.T) {
	original := []int{23, 53, 12, 44, 0}
	expected := []int{0, 12, 23, 44, 53}
	result := bubbleSort(original)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result = %v, expected = %v", result, expected)
	}
}
