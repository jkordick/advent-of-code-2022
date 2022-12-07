package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const HOUNDREDTHOUSAND = 100000

var current_directory string = "/"
var above_directory string

func main() {

	input_data, err := ioutil.ReadFile("input.txt")

	input_data_test, err := ioutil.ReadFile("test_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")
	input_data_test_string_array := strings.Split(string(input_data_test), "\n")
	fmt.Println(input_data_string_array[0])
	fmt.Println(input_data_test_string_array)

	findFiles(input_data_test_string_array)
	for i := 0; i < len(input_data_test_string_array); i++ {

	}
}

var files []int

func findFiles(str []string) {
	for i := 0; i < len(str); i++ {
		if isFile(str[i]) {
			fmt.Println(str[i], " is a file")
		} else {
			fmt.Println(str[i], " is not a file")
		}
	}

}

// func parseCommand(command string) (string, string, string) {
// 	//command_array := strings.Split(command, " ")
// }

func findDirectory(directory string) {

}

func isCommand(str string) bool {
	if strings.Contains(str, "$") {
		return true
	}
	return false
}

func isFile(str string) bool {
	file_array := strings.Split(str, " ")
	size, err := strconv.Atoi(file_array[0])
	if err != nil {
		return false
	}

	fmt.Println(size)
	return true
}

func isDirectory(str string) bool {
	if strings.Contains(str, "dir") {
		return true
	}
	return false
}

// find directories: cd dir + ls + optional(if subdirectory: dir name) + [files 0-n]

// need to know in which directory I am and which directory is above
