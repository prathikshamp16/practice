package main

import (
	"fmt"
	"net/http"
	"strconv"

	 "github.com/gin-gonic/gin"
)

func prathiksha() {
	fmt.Println("Pratii kshanaa kandaga ninnanuu adeno eno eno agideeee")
}
func main() {
	r := gin.Default()

	// Route
	r.GET("/check/:number", func(c *gin.Context) {
		numStr := c.Param("number")

		num, err := strconv.Atoi(numStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please provide a valid number",
			})
			return
		}

		if num%2 == 0 {
			c.JSON(http.StatusOK, gin.H{
				"number": num,
				"result": "Even",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"number": num,
				"result": "Odd",
			})
		}
	})

	r.Run(":8080") // start server
}
