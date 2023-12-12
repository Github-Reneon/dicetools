package dicetools

import (
	"fmt"
)

func RollNotation(command string) int {
	// Check that the format is correct
	if isCorrectNotation(&command) {
		// Return the rolled result
		return receiveRollResult(&command)
	} else {
		// Error
		fmt.Println("Incorrect notation: ", command)
		return -1
	}
}

func CheckIsCorrectNotation(command string) bool {
	return isCorrectNotation(&command)
}
