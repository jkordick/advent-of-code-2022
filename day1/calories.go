package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var total_calories []int

func main() {
	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n\n")

	for i := 0; i < len(input_data_string_array); i++ {
		needed_array := strings.Split(input_data_string_array[i], "\n")
		total_calories = append(total_calories, addUpCalories(convertStringToIntArray(needed_array)))
	}

	sort.Ints(total_calories)

	fmt.Println("highest calories:", total_calories[len(total_calories)-1])

}

func addUpCalories(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func convertStringToIntArray(arr []string) []int {
	var int_array []int
	for i := 0; i < len(arr); i++ {
		s, err := strconv.Atoi(arr[i])
		if err != nil {
			log.Fatal(err)
		}
		int_array = append(int_array, s)
	}
	return int_array
}
