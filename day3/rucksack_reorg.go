package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var priority_map = make(map[string]int)
var priority_value int

func main() {
	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")

	setupPriorityMap()

	for i := 0; i < len(input_data_string_array); i++ {
		firstHalf, secondHalf := halveString(input_data_string_array[i])
		priority_char := compareChars(firstHalf, secondHalf)
		priority_value += assignPriority(priority_char)
	}

	fmt.Println("summed up priorities:", priority_value)

}

func halveString(str string) (string, string) {
	return str[:len(str)/2], str[len(str)/2:]
}

func compareChars(str1, str2 string) string {
	for _, char1 := range str1 {
		for _, char2 := range str2 {
			if string(char1) == string(char2) {
				return string(char1)
			}
		}
	}
	return ""
}

func assignPriority(char string) int {
	return priority_map[char]
}

// a - z -> 1 - 26, A - Z -> 27 - 52
func setupPriorityMap() {
	for i := 'a'; i <= 'z'; i++ {
		priority_map[string(i)] = int(i - 'a' + 1)
	}

	for i := 'A'; i <= 'Z'; i++ {
		priority_map[string(i)] = int(i - 'A' + 27)
	}
}
