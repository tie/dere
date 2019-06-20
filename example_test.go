package dere_test

import (
	"fmt"

	"github.com/tie/dere"
)

func Example() {
	type Z struct {
		SlicePointer *[]string
		Pointer      *string
		Value        string
		Slice        []string
		Chan         chan string
		Map          map[string]string

		StructValue struct {
			Cycle *Z
		}
		StructPointer *struct {
			Cycle Z
		}
	}

	// DeepZero returns nil if argument is nil interface.
	if dere.DeepZero(nil) != nil {
		return
	}

	c, ok := dere.DeepZero(Z{
		Value: "non-empty",
	}).(Z)
	if !ok {
		return
	}
	if c.Pointer == nil || c.SlicePointer == nil {
		return
	}

	fmt.Printf("slice pointer: &%#v\n", *c.SlicePointer)
	fmt.Printf("pointer: &%#v\n", *c.Pointer)
	fmt.Printf("value: %#v\n", c.Value)
	fmt.Printf("slice: %#v\n", c.Slice)
	fmt.Printf("chan: %#v\n", c.Chan)
	fmt.Printf("map: %#v\n", c.Map)
	fmt.Printf("struct value cycle: %v\n", c.StructValue.Cycle)
	fmt.Printf("struct pointer cycle: %v\n", c.StructPointer.Cycle.StructPointer)
	// Output:
	//
	// slice pointer: &[]string(nil)
	// pointer: &""
	// value: ""
	// slice: []string(nil)
	// chan: (chan string)(nil)
	// map: map[string]string{}
	// struct value cycle: <nil>
	// struct pointer cycle: <nil>
}
