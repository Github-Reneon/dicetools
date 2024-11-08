package dicetools

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func isCorrectNotation(command *string) bool {

	// Gen regex in format to check all is correct (e.g. 1d100+10)
	reNotation, err := regexp.Compile(`^(\d+d\d+)((\+|\-)\d+d\d+)*(\+\d+|\-\d+)*$`)
	//reNotation, err := regexp.Compile(`^[1-9]\d+d[1-9]\d+$`)

	// Check for error
	if err != nil {
		fmt.Println("Error compiling regex: ", err)
		return false
	}

	//Return the result of the matching string
	return reNotation.MatchString(*command)
}

func receiveRollResult(command *string, r *rand.Rand) int {
	pattern := `[\+\-]`
	total := 0

	reClauses, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	clauses := reClauses.Split(*command, -1)

	// get the int value of all of the clauses
	parsedValues := make([]int, len(clauses))

	for clausePos, clause := range clauses {
		if strings.Contains(clause, "d") {
			diceSplit := strings.Split(clause, "d")
			numOfDice, err := strconv.Atoi(diceSplit[0])
			if err != nil {
				panic(err)
			}
			maxRoll, err2 := strconv.Atoi(diceSplit[1])
			if err2 != nil {
				panic(err2)
			}
			diceVal := 0
			for i := 0; i < numOfDice; i++ {
				diceVal += r.Intn(maxRoll) + 1
			}
			parsedValues[clausePos] = diceVal
		} else {
			parsedValues[clausePos], err = strconv.Atoi(clause)
			if err != nil {
				panic(err)
			}
		}
	}

	plusMinus := GetPlusMinus(command)

	total = parsedValues[0]

	for i, runeval := range plusMinus {
		if runeval == '+' {
			total += parsedValues[i+1]
		} else {
			total -= parsedValues[i+1]
		}
	}

	return total
}

func GetPlusMinus(command *string) string {
	ret := ""

	rePlusMinus, err := regexp.Compile(`[\+\-]`)
	if err != nil {
		panic(err)
	}

	matches := rePlusMinus.FindAllString(*command, -1)

	for _, match := range matches {
		ret += match
	}

	return ret
}
