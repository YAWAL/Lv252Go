package main

import "testing"

func Test1SearchOneRoot(t *testing.T)  {
	inputValues := []int{5, 8, 18, 41,57677}
	for _, value := range inputValues{
		if SearchOneRoot(value){
			t.Log("Function works properly with correct n-values")
		}
	}
}

func Test2SearchOneRoot(t *testing.T)  {
	inputValues := []int{3, 4, 7, 9, 11}
	for _, value := range inputValues{
		if !SearchOneRoot(value){
			t.Log("Function works properly with not correct n-values")
		}
	}
}
