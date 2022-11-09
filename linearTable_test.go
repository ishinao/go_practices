package main

import (
	"reflect"
	"testing"
)

func TestLinearTable1(t *testing.T) {
	original := []int{90, 70, 50, 80, 60, 85}
	expected := []int{90, 70, 50, 80, 60, 85}
	result := linearTable(original)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("original = %v, result= %v", original, result)
	}
}

func TestLinearTable2(t *testing.T) {
	original := []int{90, 70}
	expected := []int{90, 70}
	result := linearTable(original)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("original = %v, result= %v", original, result)
	}
}
