package main

import "testing"

//TestSearchOneRoot check if function gives root or nothing in case of root is absent
func TestSearchOneRoot(t *testing.T)  {

	inputValues := []int{5, 41,57677}

	if SearchOneRoot(5){
		t.Log("Function works properly")
	}
}
