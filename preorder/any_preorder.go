package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compare(x, y interface{}) bool {
	switch x := x.(type) {
	case int:
		if y, ok := y.(int); ok {
			return x <= y
		}
	case string:
		if y, ok := y.(string); ok {
			return x <= y
		}
	}
	return false
}

func generatePreorderRelation(set []interface{}) map[[2]interface{}]bool {
	preorderRelation := make(map[[2]interface{}]bool)

	for _, element := range set {
		preorderRelation[[2]interface{}{element, element}] = true

		for _, otherElement := range set {
			if compare(element, otherElement) {
				preorderRelation[[2]interface{}{element, otherElement}] = true
			}
		}
	}

	return preorderRelation
}

func main() {
	fmt.Println("Enter a list of elements separated by space (e.g., integers or strings):")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputElements := strings.Fields(input)
	var set []interface{}
	for _, elemStr := range inputElements {
		if num, err := strconv.Atoi(elemStr); err == nil {
			set = append(set, num)
		} else {
			set = append(set, elemStr)
		}
	}

	preorder := generatePreorderRelation(set)

	for relation, isRelated := range preorder {
		fmt.Printf("%v: %t\n", relation, isRelated)
	}
}
