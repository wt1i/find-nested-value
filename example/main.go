package main

import (
	finder "find_nested_value"
	"fmt"
)

type DSData interface {
	GetValue() string
}

func (d D) GetValue() string {
	return "test"
}

type D struct {
	Value string
}

type C struct {
	Keys map[string]DSData
}

type B struct {
	C *C
}

type A struct {
	B *B
}

func main() {
	keys := map[string]any{
		"k1": A{
			B: &B{
				C: &C{
					Keys: map[string]DSData{"k2": D{
						Value: "101",
					}}},
			},
		},
	}

	value, err := finder.FindNestedValue(keys, "k1.B.C.Keys.k2.Value")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("value from type:%T value: %v\n", value, value)
	}
}
