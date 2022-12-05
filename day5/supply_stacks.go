package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var stack [][]string

func main() {
	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")

	commands := input_data_string_array[10:]

	base_stack := input_data_string_array[:8]

	parseCargo(base_stack, commands)

}

func parseCargo(b_stack []string, commands []string) {
	for i := 0; i < len(b_stack); i++ {
		b_stack[i] = strings.Replace(b_stack[i], "    ", "[0]", -1)
		b_stack[i] = strings.Replace(b_stack[i], " ", "", -1)
		b_stack[i] = strings.Replace(b_stack[i], "[", "", -1)
		parsed_stack := strings.Split(b_stack[i], "]")

		for index, value := range parsed_stack {
			var l []string
			if len(stack) <= index {
				l = make([]string, 0)
				stack = append(stack, l)
			} else {
				l = stack[index]
			}
			if value != "0" {
				l = append(l, value)
				stack[index] = l
			}
		}

		moveCargo(stack, commands)
	}

	// for _, value := range stack {
	// 	fmt.Println(value[0])
	// }
}

func moveCargo(stack [][]string, commands []string) {
	commands[i] = strings.Replace(commands[i], "move ", "", -1)
	commands[i] = strings.Replace(commands[i], " from ", ",", -1)
	commands[i] = strings.Replace(commands[i], " to ", ",", -1)
	parsed_commands := strings.Split(commands[i], ",")

	elements, _ := strconv.Atoi(parsed_commands[0])
	//from, _ := strconv.Atoi(parsed_commands[1])
	//to, _ := strconv.Atoi(parsed_commands[2])

	for range makeRange(1, elements) {
		// v := stack[from-1][0]
		// l := append([]string{v}, stack[to-1]...)
		// stack[to-1] = l
		// stack[from-1] = stack[from-1][1:]
	}

}

func makeRange(min, max int) []int {
	r := make([]int, max-min+1)
	for i := range r {
		r[i] = min + i
	}
	return r
}
