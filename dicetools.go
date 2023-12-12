package dicetools

import (
	"fmt"
	"regexp"
)

func IsCorrectNotation(command *string) bool {

	// Gen regex in format to check all is correct (e.g. 1d100+10)
	reNotation, err := regexp.Compile(`^\d+d\d+(\+\d+)?(\-\d+)?$`)

	// Check for error
	if err != nil {
		fmt.Println("Error compiling regex: ", err)
		return false
	}

	//Return the result of the matching string
	return reNotation.MatchString(*command)
}

func ReceiveRollResult(command *string) int {
	/*
		1. Split the string into the three possible clauses
			a. Split by +
			b. Split by -
		2. Split the first clause into the number of rolls and the maximum of the roll
		3. Add the second clause
		4. Subtract the third clause
	*/
	return 0
	//if strings.Contains(*command, '+')
}

func RollNotation(command string) int {
	// Check that the format is correct
	if IsCorrectNotation(&command) {
		// Return the rolled result
		return ReceiveRollResult(&command)
	} else {
		// Error
		fmt.Println("Incorrect notation: ", command)
		return 0
	}
}
