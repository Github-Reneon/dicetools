package dicetools

import (
	"fmt"
	"regexp"
)

func isCorrectNotation(command *string) bool {

	// Gen regex in format to check all is correct (e.g. 1d100+10)
	reNotation, err := regexp.Compile(`^([1-9]\d+d[1-9]\d+)*(\+\d+|\-\d+)*$`)

	// Check for error
	if err != nil {
		fmt.Println("Error compiling regex: ", err)
		return false
	}

	//Return the result of the matching string
	return reNotation.MatchString(*command)
}

func receiveRollResult(command *string) int {
	/*
		1. Split the string into the three possible clauses
			a. Split by +
			b. Split by -
		2. Split the first clause into the number of rolls and the maximum of the roll
		3. Add the second clause
		4. Subtract the third clause
	*/

	pattern := `[\+-]`
	total := 0

	reClauses, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	result := reClauses.Split(*command, -1)

	fmt.Println(result)

	return total
}
