package main

import (
	"fmt"
	"testing"
)

//条件语句

// if
func TestIf(t *testing.T) {

	var a = 2

	if a == 1 {
		fmt.Println("a=1")
	} else if a == 2 {
		fmt.Println("a=2")
	} else {
		fmt.Println("other")
	}

}

func TestSwitch(t *testing.T) {
	var a = 5
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
		break
	default:
		fmt.Println("default")
	}

}

// todo select
