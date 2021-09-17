package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokhantamkoc/quiz-app-api/internal/common"
	"github.com/gokhantamkoc/quiz-app-api/internal/quiz"
)

func main() {
	r := gin.Default()

	r.GET("/", common.CORSMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Sanctuary!",
		})
	})
	quiz.Routes(r)
	r.Run("0.0.0.0:9001")
}
