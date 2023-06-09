// Package forth implements the "Forth" exercise.
package forth

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrInvalidStack = errors.New("invalid stack")
	userDef         = regexp.MustCompile(": ([^ ]+) (.+) ;")
)

// : ([^ ]+) (.+) ;
// this regular expression will match a string that has a : character, followed by one or more non-space characters,
// followed by a space, then any characters (one or more), another space and ends with a ; character.

func Forth(inputData []string) ([]int, error) {

	customMappings := map[string]string{}

	// Detect custom defined variables
	for index := 0; index < len(inputData)-1; index++ {
		matchGroups := userDef.FindStringSubmatch(strings.ToUpper(inputData[index]))
		if len(matchGroups) == 3 {
			for key, value := range customMappings {
				matchGroups[2] = strings.ReplaceAll(matchGroups[2], key, value)
			}
			customMappings[matchGroups[1]] = matchGroups[2]
		}
	}

	// Generate initial string with replaced user defined variables
	baseString := strings.ToUpper(inputData[len(inputData)-1])
	for key, value := range customMappings {
		baseString = strings.ReplaceAll(baseString, key, value)
	}

	// Initialize the stack
	processingStack := make([]int, 0, 2)

	// Loop through the string to fill the stack
	for _, item := range strings.Split(baseString, " ") {
		switch item {
		case "+":
			if len(processingStack) < 2 {
				return nil, ErrInvalidStack
			}
			processingStack[len(processingStack)-2] += processingStack[len(processingStack)-1]
			processingStack = processingStack[:len(processingStack)-1]
		case "-":
			if len(processingStack) < 2 {
				return nil, ErrInvalidStack
			}
			processingStack[len(processingStack)-2] -= processingStack[len(processingStack)-1]
			processingStack = processingStack[:len(processingStack)-1]
		case "*":
			if len(processingStack) < 2 {
				return nil, ErrInvalidStack
			}
			processingStack[len(processingStack)-2] *= processingStack[len(processingStack)-1]
			processingStack = processingStack[:len(processingStack)-1]
		case "/":
			if len(processingStack) < 2 || processingStack[len(processingStack)-1] == 0 {
				return nil, ErrInvalidStack
			}
			processingStack[len(processingStack)-2] /= processingStack[len(processingStack)-1]
			processingStack = processingStack[:len(processingStack)-1]
		case "DUP":
			if len(processingStack) == 0 {
				return nil, ErrInvalidStack
			}
			processingStack = append(processingStack, processingStack[len(processingStack)-1])
		case "DROP":
			if len(processingStack) == 0 {
				return nil, ErrInvalidStack
			}
			processingStack = processingStack[:len(processingStack)-1]
		case "SWAP":
			if len(processingStack) < 2 {
				return nil, ErrInvalidStack
			}
			processingStack[len(processingStack)-2], processingStack[len(processingStack)-1] = processingStack[len(processingStack)-1], processingStack[len(processingStack)-2]
		case "OVER":
			if len(processingStack) < 2 {
				return nil, ErrInvalidStack
			}
			processingStack = append(processingStack, processingStack[len(processingStack)-2])
		default:
			intValue, convertErr := strconv.Atoi(item)
			if convertErr != nil {
				return nil, convertErr
			}
			processingStack = append(processingStack, intValue)
		}
	}
	return processingStack, nil
}
