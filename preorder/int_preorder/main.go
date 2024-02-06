package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generatePreorderRelation(set []int) map[[2]int]bool {
	preorderRelation := make(map[[2]int]bool)

	for _, element := range set {
		preorderRelation[[2]int{element, element}] = true

		for _, otherElement := range set {
			if element <= otherElement {
				preorderRelation[[2]int{element, otherElement}] = true
			}
		}
	}

	return preorderRelation
}

func main() {
	fmt.Println("Enter a list of integers separated by space:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	inputNumbers := strings.Fields(input)
	var set []int
	for _, numStr := range inputNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Error converting '%s' to int: %s\n", numStr, err)
			return
		}
		set = append(set, num)
	}

	preorder := generatePreorderRelation(set)

	for relation, isRelated := range preorder {
		fmt.Printf("%v: %t\n", relation, isRelated)
	}
}
