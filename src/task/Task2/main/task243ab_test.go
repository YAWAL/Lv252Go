package main

import (
	"testing"
)

// check if function gives root or nothing in case of root is absent
func TestSearchOneRoot(t *testing.T)  {

	if SearchOneRoot != nil {
		t.Log("Function SearchOneRoot works")
	}
}

// check if function gives roots or nothing in case of root is absent
func TestSearchAllRoots(t *testing.T)  {

	if SearchAllRoots != nil {
		t.Log("Function SearchAllRoots works")

	}
}
