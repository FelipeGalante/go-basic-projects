package main

import "fmt"

func main() {
	// Go has 4 different types categories
	// Basic Types
	// Aggregate Types - Arrays and structs
	// Reference Types - Pointers and slices
	// Interface Types - Standard interfaces
	integers()
}

func integers() {
	// unsigned int with 8 bits
	// Can store: 0 to 255
	var uint1 uint8
	// signed int with 8 bits
	// Can store: -127 to 127
	var int1 int8

	// unsigned int with 16 bits
	var uint2 uint16
	// signed int with 16 bits
	var int2 int16

	// unsigned int with 32 bits
	var uint3 uint32
	// signed int with 32 bits
	var int3 int32

	// unsigned int with 64 bits
	var uint4 uint64
	// signed int with 64 bits
	var int4 int64

	fmt.Printf("%T\n", uint1)
	fmt.Printf("%T\n", int1)
	fmt.Printf("%T\n", uint2)
	fmt.Printf("%T\n", int2)
	fmt.Printf("%T\n", uint3)
	fmt.Printf("%T\n", int3)
	fmt.Printf("%T\n", uint4)
	fmt.Printf("%T\n", int4)
}
