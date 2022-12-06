package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var first_marker_counter int

const FOUR int = 4
const FOURTEEN int = 14

func main() {

	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")
	input_data_string := strings.Join(input_data_string_array, "")

	checkDuplicates(input_data_string, FOUR)

	checkDuplicates(input_data_string, FOURTEEN)
}

func checkDuplicates(input string, counter int) {
	var no_duplicate int
	for i := 0; i < len(input); i++ {
		current_list := input[i : i+counter]

		if no_duplicate != counter {
			if strings.Contains(current_list[1:], string(input[i])) {
				no_duplicate = 0
			} else {
				no_duplicate++
			}

		} else {
			first_marker_counter = i
			fmt.Println("++++ found", counter, " no duplicates ++++++", first_marker_counter)
			break
		}
	}

}

func checkDuplicates14(input string) {
	var no_duplicate int
	var current_list string
	//var yann_ist_schlau int
	for i := 0; i < len(input); i++ {

		if (i + FOURTEEN) > len(input) {
			current_list = input[i:]
			//yann_ist_schlau++
		} else {
			current_list = input[i : i+FOURTEEN]
		}

		//fmt.Println(current_list)

		if no_duplicate != FOURTEEN { //-(len(input)-i-FOURTEEN
			if strings.Contains(current_list[1:], string(input[i])) {
				no_duplicate = 0
			} else {
				no_duplicate++
			}

		} else {
			first_marker_counter = i
			fmt.Println("++++ found FOURTEEN no duplicates ++++++", first_marker_counter)
			break
		}
	}

}
