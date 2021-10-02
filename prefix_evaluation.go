package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

// eval evaluates "arg1 op arg2"
func eval(op string, arg1, arg2 float64) (float64, error) {
	switch op {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "^":
		return math.Pow(arg1, arg2), nil
	case "/":
		if arg2 == 0 {
			return 0, fmt.Errorf("divided by zero: %s %f %f", op, arg1, arg2)
		}
		return float64(arg1) / float64(arg2), nil
	default:
		return 0, fmt.Errorf("unsupported operator - %s %f %f", op, arg1, arg2)
	}
}
func Prefix(r io.Reader) (ret float64, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	if !scanner.Scan() {
		return
	}
	return prefix(scanner.Text(), scanner)
}
func prefix(op string, scanner *bufio.Scanner) (float64, error) {
	if !scanner.Scan() {
		// no more tokens - this is allowed if the given operator is actually a number
		if val, err := strconv.ParseFloat(op, 64); err == nil {
			return val, nil
		} else {
			return 0, fmt.Errorf(`last token is not a number: %v`, err)
		}
	}

	// Evaluate the first argument
	arg1 := scanner.Text()
	val1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		// token arg1 is an operator
		if val1, err = prefix(arg1, scanner); err != nil {
			return 0, err
		}
	}

	// Evaluate the second argument
	if !scanner.Scan() {
		// no more token
		return 0, fmt.Errorf(`missing second parameter for eval %s`, op)
	}
	arg2 := scanner.Text()
	val2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		if val2, err = prefix(arg2, scanner); err != nil {
			return 0, err
		}
	}
	return eval(op, val1, val2)
}
