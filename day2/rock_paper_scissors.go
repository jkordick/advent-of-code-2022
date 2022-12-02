package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/**
* Part 1
* A X Rock 1
* B Y Paper 2
* C Z Scissors 3
* A > Z
* B > X
* C > Y
* Lost 0
* Draw 3
* Win 6
 */

func main() {
	input_data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input_data_string_array := strings.Split(string(input_data), "\n")
	var pointsByShape int
	var pointsAsPlanned int

	for i := 0; i < len(input_data_string_array); i++ {
		needed_array := strings.Split(input_data_string_array[i], " ")
		pointsByShape += calculatePointsByShape(needed_array)
		pointsAsPlanned += calculatePointsAsPlanned(needed_array)
	}

	fmt.Println("summed up points by shape:", pointsByShape)
	fmt.Println("summed up points as planned:", pointsAsPlanned)
}

const rock int = 1
const paper int = 2
const scissors int = 3

const lose int = 0
const draw int = 3
const win int = 6

/*
* Part 2
* X Lose
* Y Draw
* Z Win
 */

func calculatePointsAsPlanned(arr []string) int {
	temp := 0

	if arr[0] == "A" { // Opponent choose Rock
		if arr[1] == "Z" { // Win with Paper
			temp += win + paper
		} else if arr[1] == "X" { // Lose with scissors
			temp += lose + scissors
		} else if arr[1] == "Y" { // Draw with rock
			temp += draw + rock
		}
	} else if arr[0] == "B" { //  Opponent choose Paper
		if arr[1] == "Z" { // Win with Scissors
			temp += win + scissors
		} else if arr[1] == "X" { // Lose with Rock
			temp += lose + rock
		} else if arr[1] == "Y" { // Draw with Paper
			temp += draw + paper
		}
	} else if arr[0] == "C" { //  Opponent choose Scissors
		if arr[1] == "Z" { // Win with rock
			temp += win + rock
		} else if arr[1] == "X" { // Lose with Paper
			temp += lose + paper
		} else if arr[1] == "Y" { // Draw with Scissors
			temp += draw + scissors
		}
	}

	return temp
}

func calculatePointsByShape(arr []string) int {
	temp := 0

	if arr[0] == "A" { // Opponent choose Rock
		if arr[1] == "Z" { // Against Scissors lost
			temp += lose + scissors
		} else if arr[1] == "X" { // Against Rock draw
			temp += draw + rock
		} else if arr[1] == "Y" { // Against Paper win
			temp += win + paper
		}
	} else if arr[0] == "B" { //  Opponent choose Paper
		if arr[1] == "Z" { // Against Scissors win
			temp += win + scissors
		} else if arr[1] == "X" { // Against Rock lost
			temp += lose + rock
		} else if arr[1] == "Y" { // Against Paper draw
			temp += draw + paper
		}
	} else if arr[0] == "C" { //  Opponent choose Scissors
		if arr[1] == "Z" { // Against Scissors draw
			temp += draw + scissors
		} else if arr[1] == "X" { // Against Rock win
			temp += win + rock
		} else if arr[1] == "Y" { // Against Paper lost
			temp += lose + paper
		}
	}

	return temp
}
