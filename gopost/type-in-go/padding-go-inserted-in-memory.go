package main

import "unsafe"

type Example struct {
	BoolValue  bool
	IntValue   int16
	FloatValue float32
}

func main() {
	example := &Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	exampleNext := &Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	alignmentBoundary := unsafe.Alignof(example)

	sizeBool := unsafe.Sizeof(example.BoolValue)
	offsetBool := unsafe.Offsetof(example.BoolValue)

	sizeInt := unsafe.Sizeof(example.IntValue)
	offsetInt := unsafe.Offsetof(example.IntValue)

	sizeFloat := unsafe.Sizeof(example.FloatValue)
	offsetFloat := unsafe.Offsetof(example.FloatValue)
}
