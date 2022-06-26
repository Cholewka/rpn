package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

var precedence map[string]int
var associatives map[string]bool

func init() {
	precedence = make(map[string]int)
	associatives = make(map[string]bool)

	precedence["+"] = 1
	precedence["-"] = 1
	precedence["/"] = 2
	precedence["*"] = 2
	precedence["^"] = 3

	associatives["^"] = true
}

func prepend(arr []string, str string) []string {
	arr = append(arr, "")
	copy(arr[1:], arr)
	arr[0] = str
	return arr
}

func pop(arr []string) []string {
	arr = arr[1:]
	return arr
}

func main() {
	functions := []string{"sin", "cos", "tg", "ctg"}

	reader := bufio.NewReader(os.Stdin)
	var expr string

	for {
		fmt.Print("enter expression: ")

		e, err := reader.ReadString('\n')
		if err != nil {
			panic("cannot read from standard input", false)
			continue
		}

		if e = strings.TrimSpace(e); len(e) == 0 {
			panic("expression is empty", false)
			continue
		}

		expr = e
		break
	}

	outputQueue := []string{}
	operatorStack := []string{}
	tokens := scanExpr(expr)
	fmt.Println(tokens)

	for _, t := range tokens {
		if _, err := strconv.Atoi(t); err == nil {
			outputQueue = append(outputQueue, t)
		} else if contains(functions, t) {
			operatorStack = prepend(operatorStack, t)
		} else if precedence[t] > 0 {
			for {
				if len(operatorStack) > 0 {
					o2 := operatorStack[0]

					if len(o2) > 0 && o2 != "(" && (precedence[o2] > precedence[t] || precedence[o2] == precedence[t] && !associatives[o2]) {
						operatorStack = pop(operatorStack)
						outputQueue = append(outputQueue, o2)
					} else {
						break
					}
				} else {
					break
				}
			}

			operatorStack = prepend(operatorStack, t)
		} else if t == "(" {
			operatorStack = prepend(operatorStack, t)
		} else if t == ")" {
			for {
				if len(operatorStack) > 0 && operatorStack[0] != "(" {
					o2 := operatorStack[0]
					operatorStack = pop(operatorStack)
					outputQueue = append(outputQueue, o2)
				} else {
					break
				}
			}

			if len(operatorStack) > 0 && operatorStack[0] == "(" {
				operatorStack = pop(operatorStack)

				if len(operatorStack) > 0 && contains(functions, operatorStack[0]) {
					fn := operatorStack[0]
					operatorStack = pop(operatorStack)
					outputQueue = append(outputQueue, fn)
				}
			}
		}

		fmt.Printf("%q\t%v\t%v\n", t, outputQueue, operatorStack)
	}

	for _, o := range operatorStack {
		if o == "(" {
			panic("invalid expression", true)
		}

		operatorStack = operatorStack[:len(operatorStack)-1]
		outputQueue = append(outputQueue, o)
	}

	fmt.Printf("\nRESULT : %v\n", outputQueue)
}

func scanExpr(expr string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(expr))

	var token rune
	var result []string

	for token != scanner.EOF {
		token = s.Scan()
		value := strings.TrimSpace(s.TokenText())

		if len(value) > 0 {
			result = append(result, s.TokenText())
		}
	}

	return result
}

func panic(msg string, exit bool) {
	fmt.Fprintf(os.Stderr, "rpn: %s\n", msg)

	if exit {
		os.Exit(1)
	}
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}
