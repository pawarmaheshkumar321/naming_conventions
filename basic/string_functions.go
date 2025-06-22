package main

import (
	"fmt"
	"strings"
)

var builder1 strings.Builder

func main() {
	builder1.WriteString("Hello")
	builder1.WriteString(", ")
	builder1.WriteString("World")

	result := builder1.String()

	fmt.Println(result)

	builder1.WriteRune('c')

	result = builder1.String()

	fmt.Println(result)

	builder1.Reset()
}
