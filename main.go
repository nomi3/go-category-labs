package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generatePreorderRelation(set []int) map[string]bool {
	preorderRelation := make(map[string]bool)
	for _, element := range set {
		preorderRelation[fmt.Sprintf("[%d,%d]", element, element)] = true
		for _, otherElement := range set {
			if element <= otherElement {
				preorderRelation[fmt.Sprintf("[%d,%d]", element, otherElement)] = true
			}
		}
	}
	return preorderRelation
}

func main() {
	r := gin.Default()

	r.POST("/preorder", func(c *gin.Context) {
		var set []int
		if err := c.BindJSON(&set); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		preorderRelation := generatePreorderRelation(set)
		c.JSON(http.StatusOK, preorderRelation)
	})

	r.Run()
}
