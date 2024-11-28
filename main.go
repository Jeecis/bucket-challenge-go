package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateBucket(bucketArray []int) int {
	bestValue := 0
	arrayLength := len(bucketArray)

	endPosition := arrayLength - 1
	var endPointer int

	startPosition := 0
	var startPointer int

	for i := 0; i < len(bucketArray); i++ {
		startPointer = bucketArray[startPosition]
		endPointer = bucketArray[endPosition]
		xLength := endPosition - startPosition
		var yLength int

		if endPointer > startPointer {
			yLength = startPointer
			startPosition++

		} else {
			yLength = endPointer
			endPosition--
		}

		if xLength*yLength > bestValue {
			bestValue = xLength * yLength
		}
	}

	return bestValue
}

func readValues(path string) ([]int, error) {
	byteData, err := os.ReadFile(path) // Replace with your file path
	if err != nil {
		return nil, err
	}

	strData := string(byteData)
	stringArray := strings.Split(strData, "\n")

	var valueArray []int

	for _, value := range stringArray {
		num, err := strconv.Atoi(value) // Convert string to int
		if err != nil {
			return nil, err
		}
		valueArray = append(valueArray, num)
	}

	return valueArray, nil
}

func main() {

	path := flag.String("path", "", "Path to the txt file")
	flag.Parse()

	if path == nil {
		fmt.Println("usage: -path path to your file")
		return
	}

	values, err := readValues(*path)
	if err != nil {
		fmt.Println(err)
	}

	bestValue := calculateBucket(values)
	fmt.Println(bestValue)

}
