package main

import "fmt"

// Infix to prefix conversion

// Stack node
type StackNode struct {
	// Stack data
	element string
	next    *StackNode
}

func getStackNode(element string,
	next *StackNode) *StackNode {
	var me *StackNode = &StackNode{}
	me.element = element
	me.next = next
	return me
}

// Define a custom stack
type MyStack struct {
	top  *StackNode
	size int
}

func getMyStack() *MyStack {
	// Set node values
	return &MyStack{nil, 0}
}

// Add node at the top of stack
func (this *MyStack) push(element string) {
	this.top = getStackNode(element, this.top)
	this.size++
}
func (this MyStack) isEmpty() bool {
	if this.size > 0 && this.top != nil {
		return false
	} else {
		return true
	}
}

// Remove top element of stack
func (this *MyStack) pop() string {
	var c string = ""
	if this.size > 0 && this.top != nil {
		var temp *StackNode = this.top
		// Change top element of stack
		this.top = temp.next
		this.size--
		c = temp.element
	}
	return c + " "
}

// Return top element of stack
func (this MyStack) peek() string {
	if this.top == nil {
		return ""
	}
	return this.top.element
}

type Conversion struct{}

func getConversion() *Conversion {
	return &Conversion{}
}
func (this Conversion) precedence(text string) int {
	if text == "+" || text == "-" {
		return 1
	} else if text == "*" || text == "/" {
		return 2
	} else if text == "^" {
		return 3
	}
	return -1
}
func (this Conversion) isOperator(text string) bool {
	if this.precedence(text) != -1 {
		return true
	}
	return false
}

// Converting the given infix
// expression to prefix expression
func (this Conversion) infixToPrefix(infix string) string {
	// Get the size
	var size int = len(infix)
	// Create stack object
	var s *MyStack = getMyStack()
	var op *MyStack = getMyStack()
	// Some useful variables which is using of to
	// storing current result
	var auxiliary string = ""
	var op1 string = ""
	var op2 string = ""
	var isValid bool = true
	for i := 0; i < size && isValid; i++ {
		if infix[i] == '(' {
			op.push("(")
		} else if this.isOperator(string(infix[i])) == true {
			// When get a operator
			for s.size > 1 && this.precedence(
				string(infix[i])) <= this.precedence(op.peek()) {
				op1 = s.pop()
				op2 = s.pop()
				auxiliary = op.pop() + op2 + op1
				// Add new result into stack
				s.push(auxiliary)
			}
			//auxiliary = string(infix[i]) + " "// print with spaces
			auxiliary = string(infix[i])
			op.push(auxiliary)
		} else if infix[i] == ')' {
			if s.size > 1 {
				for s.size > 1 && op.peek() != "(" {
					op1 = s.pop()
					op2 = s.pop()
					auxiliary = op.pop() + op2 + op1
					// Add new result into stack
					s.push(auxiliary)
				}
				op.pop()
			} else {
				isValid = false
			}
		} else if (infix[i] >= '0' && infix[i] <= '9') ||
			(infix[i] >= 'a' && infix[i] <= 'z') ||
			(infix[i] >= 'A' && infix[i] <= 'Z') {
			//auxiliary = string(infix[i]) + " " // print with spaces
			auxiliary = string(infix[i])
			s.push(auxiliary)
		} else {
			isValid = false
		}
	}
	if isValid == false {
		fmt.Print("Invalid infix : ", infix)
	} else {
		// Display result
		fmt.Printf("Infix  : %s ", infix)
		result := s.pop()
		fmt.Printf("\n Prefix : %s\n", result)
		return result
	}
	return "ERROR"
}
