package main

import (
	"testing"
)

func TestStringFilter(t *testing.T) {
	// compare empty strings
	val := ""
	if err := assertEq(true, (&Filter{Operator: Equal, StrValue: ""}).MatchString(&val)); err != nil {
		t.Error(err)
	}
}

func TestStringListFilter(t *testing.T) {
	value := []string{"abc", "def"}
	if err := assertEq(true, (&Filter{Operator: GreaterThan, StrValue: "def"}).MatchStringList(&value)); err != nil {
		t.Error(err)
	}
}

func TestIntListFilter(t *testing.T) {
	value := []int{1, 2, 3, 4, 5}
	if err := assertEq(true, (&Filter{Operator: GreaterThan, FloatValue: 5}).MatchIntList(value)); err != nil {
		t.Error(err)
	}
}
