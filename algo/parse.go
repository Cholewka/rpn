// Algo is an implementation of the shunting yard algorithm.
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
package algo

import (
	"errors"
	"strconv"
	"strings"
)

var (
	operatorStack stack
	outputQueue   []string
	functions     []string
	precedence    map[string]int
	associatives  map[string]bool
)

func init() {
	precedence = make(map[string]int)
	associatives = make(map[string]bool)

	functions = []string{"sin", "cos", "tg", "ctg"}
	precedence["+"] = 1
	precedence["-"] = 1
	precedence["*"] = 2
	precedence["/"] = 2
	precedence["^"] = 3

	associatives["^"] = true // right-associative
}

func Parse(expr string) (result []string, err error) {
	operatorStack = stack{}
	outputQueue = []string{}
	tokens := tokenise(expr)

	// read all the tokens
	for _, token := range tokens {
		err := readToken(token)

		if err != nil {
			return []string{}, err
		}
	}

	// left operators on the operator stack; top-to-bottom
	for i := len(operatorStack) - 1; i >= 0; i-- {
		op := operatorStack[i]

		if op == "(" {
			return []string{}, errors.New("mismatched parenthesis")
		}

		operatorStack.pop()
		outputQueue = append(outputQueue, op)
	}

	return outputQueue, nil
}

func readToken(t string) error {
	if _, err := strconv.Atoi(t); err == nil {
		outputQueue = append(outputQueue, t)
		return nil
	}

	if isfunction(t) {
		operatorStack.push(t)
		return nil
	}

	if precedence[t] > 0 {
		for {
			if len(operatorStack) == 0 {
				break
			}

			op := operatorStack.top()
			if op != "(" && (precedence[op] > precedence[t] || (precedence[op] == precedence[t] && !associatives[t])) {
				operatorStack.pop()
				outputQueue = append(outputQueue, op)
				continue
			}

			break
		}

		operatorStack = append(operatorStack, t)
		return nil
	}

	if t == "(" {
		operatorStack.push(t)
		return nil
	}

	if t == ")" {
		for len(operatorStack) > 0 && operatorStack.top() != "(" {
			op := operatorStack.pop()
			outputQueue = append(outputQueue, op)
		}
		if operatorStack.top() != "(" {
			return errors.New("mismatched parenthesis")
		}
		operatorStack.pop() // discard the operator
		if isfunction(operatorStack.top()) {
			f := operatorStack.pop()
			outputQueue = append(outputQueue, f)
		}

		return nil
	}

	return nil
}

func tokenise(expr string) []string {
	return strings.Split(expr, " ")
}

func isfunction(t string) bool {
	for _, f := range functions {
		if f == t {
			return true
		}
	}
	return false
}
