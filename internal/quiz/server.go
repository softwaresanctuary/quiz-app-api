package quiz

import (
	"github.com/gin-gonic/gin"
	"github.com/gokhantamkoc/quiz-app-api/api"
	"github.com/gokhantamkoc/quiz-app-api/internal/common"
)

func Routes(engine *gin.Engine) {
	qServer := &QuizServer{
		Service: NewService(),
	}
	quiz := engine.Group("/quiz", common.CORSMiddleware())
	quiz.POST("", qServer.CreateQuiz)
	quiz.PUT("", qServer.UpdateQuiz)
	quiz.GET("/:id", qServer.GetQuiz)
	quiz.GET("/list", qServer.ListQuiz)
}

type QuizServer struct {
	Service *Service
}

func (q *QuizServer) CreateQuiz(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	var newQuiz *api.Quiz
	if err := c.BindJSON(&newQuiz); err != nil {
		c.JSON(400, &api.ClientError{
			Message: "İstemci hatası oluştu! " + err.Error(),
		})
		return
	}
	quiz, err := q.Service.CreateQuiz(newQuiz)
	if err != nil {
		c.JSON(400, &api.ClientError{
			Message: "İstemci hatası oluştu!",
		})
		return
	}
	c.JSON(200, quiz)
}

func (q *QuizServer) UpdateQuiz(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	quiz, err := q.Service.UpdateQuiz()
	if err != nil {
		c.JSON(400, &api.ClientError{
			Message: "İstemci hatası oluştu!",
		})
		return
	}
	c.JSON(200, quiz)
}

func (q *QuizServer) GetQuiz(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	quizId := c.Param("id")
	quiz, err := q.Service.GetQuiz(quizId)
	if err != nil {
		c.JSON(400, &api.ClientError{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, quiz)
}

func (q *QuizServer) ListQuiz(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, q.Service.ListQuiz())
}

