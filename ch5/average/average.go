package ch5_average

import (
	"fmt"
	"github.com/sebastiankennedy/head-first-go/ch5/datafile"
	"log"
)

func Average() {
	numbers, err := ch5_datafile.GetFloats("./data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
