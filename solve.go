// This file solves parsed queues in RPN manner.
package rpn

import (
	"errors"
	"fmt"
	"strconv"
)

func Solve(tokens []string) (result int, err error) {
	operands := stack{}

	for _, t := range tokens {
		if _, err := strconv.Atoi(t); err == nil {
			operands.push(t)
			continue
		}

		if precedence[t] > 0 {
			if len(operands) < 2 {
				return 0, errors.New("mismatched operands")
			}

			o1, o2 := operands.pop(), operands.pop()

			num1, err := strconv.Atoi(o1)
			if err != nil {
				return 0, errors.New("internal error: operand is not a digit")
			}

			num2, err := strconv.Atoi(o2)
			if err != nil {
				return 0, errors.New("internal error: operand is not a digit")
			}

			switch t {
			case "+":
				operands.push(fmt.Sprint(num1 + num2))
			case "-":
				operands.push(fmt.Sprint(num1 - num2))
			case "*":
				operands.push(fmt.Sprint(num1 * num2))
			case "/":
				operands.push(fmt.Sprint(num1 / num2))
			case "^":
				operands.push(fmt.Sprint(num1 ^ num2))
			default:
				return 0, errors.New("internal error: unknown operator")
			}

			continue
		}
	}

	if len(operands) == 0 || len(operands) > 1 {
		return 0, errors.New("internal error: unknown result")
	}

	result, err = strconv.Atoi(operands.top())
	if err != nil {
		return 0, errors.New("internal error: result is not a digit")
	}

	return result, nil
}
