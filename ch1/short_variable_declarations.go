package ch1

import "fmt"

func ShortVariableDeclarations() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Sebastian Lu"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
