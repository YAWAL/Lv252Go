package main

import (
	"testing"
)

func TestCreateSliceOfK(t *testing.T)  {
	expectedSlice := []int{0,0,0,0,0}
	createdSlice := CreateSliceOfK(5)
	for i := 0; i < len(createdSlice); i++{
		if expectedSlice[i] == createdSlice[i] {
			t.Log("Element", i, "of expected and created slices are the same")
		}
	}
}

func TestFindMaxValueOfK(t *testing.T) {
	slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	maxValueOfSlice := FindMaxValueOfK(slice)
	maxValue := 9
	if maxValueOfSlice == maxValue{
		t.Log(" Test pass")
	}
	if maxValueOfSlice != 9{
		 t.Error("Smtg wrong")
	}
}
