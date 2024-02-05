package main

import (
	"fmt"
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
	exampleSet := []int{1, 3, 5, 7}

	preorder := generatePreorderRelation(exampleSet)

	for relation, isRelated := range preorder {
		fmt.Printf("%v: %t\n", relation, isRelated)
	}
}
