package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var number_of_fully_contained_ranges int
var number_of_partially_contained_ranges int

func main() {
	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")

	for i := 0; i < len(input_data_string_array); i++ {
		sectionA, sectionB := separateSections(input_data_string_array[i])
		compareSectionRanges(sectionA, sectionB)
	}

	fmt.Println("number of fully contained ranges:", number_of_fully_contained_ranges)
	fmt.Println("number of partially contained ranges:", number_of_partially_contained_ranges)

}

func separateSections(str string) (string, string) {
	sections := strings.Split(string(str), ",")
	return sections[0], sections[1]
}

func compareSectionRanges(sectionA, sectionB string) {
	sectionA_int := convvertStringArrayToIntArray(strings.Split(sectionA, "-"))
	sectionB_int := convvertStringArrayToIntArray(strings.Split(sectionB, "-"))

	if sectionA_int[0] <= sectionB_int[0] && sectionA_int[1] >= sectionB_int[1] || sectionB_int[0] <= sectionA_int[0] && sectionB_int[1] >= sectionA_int[1] {
		number_of_fully_contained_ranges++
	}

	rangeA := makeRange(sectionA_int[0], sectionA_int[1])
	rangeB := makeRange(sectionB_int[0], sectionB_int[1])

out:
	for _, i := range rangeA {
		for _, j := range rangeB {
			if i == j {
				number_of_partially_contained_ranges++
				break out
			}
		}
	}
}

func convvertStringArrayToIntArray(str []string) []int {
	var int_array []int
	for _, i := range str {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		int_array = append(int_array, j)
	}
	return int_array
}

func makeRange(min, max int) []int {
	r := make([]int, max-min+1)
	for i := range r {
		r[i] = min + i
	}
	return r
}
