package dicetools

import (
	"math/rand"
	"fmt"
)

func RollNotation(command string, rand *rand.Rand) int {
	// Check that the format is correct
	if isCorrectNotation(&command) {
		// Return the rolled result
		return receiveRollResult(&command, rand)
	} else {
		// Error
		fmt.Println("Incorrect notation: ", command)
		return -1
	}
}

func CheckIsCorrectNotation(command string) bool {
	return isCorrectNotation(&command)
}
