package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AndrewSerra/creditcard-validator/validator"
	"github.com/gin-gonic/gin"
)

type CardRequestBody struct {
	CardNumber string `json:"card" binding:"required"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "3000")

	router := gin.Default()

	router.POST("/verify", func(c *gin.Context) {
		var body CardRequestBody

		err := c.ShouldBindJSON(&body)

		if err != nil {
			log.Println("error parsing body ", err)
			c.Status(400)
			return
		}

		result, err := validator.Validate(body.CardNumber)

		if err != nil {
			log.Println("error validating card ", err)
			c.JSON(http.StatusOK, gin.H{
				"isValid": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"isValid": result,
			"message": nil,
		})
	})

	router.Run(fmt.Sprintf(":%s", port))
}
