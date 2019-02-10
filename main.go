package main

import "fmt"

func main() {
	name := "Reggie"

	// This points to the actual memory address of "name"
	pointer := &name

	fmt.Println(&name)

	// This function will make a copy of "pointer" in another memory space and
	pointerPrint(pointer)
}

// This function will make a copy of the parameter passed and print the memory address of the pointer after it's been copied
func pointerPrint(p *string) {
	fmt.Println(&p)
}
