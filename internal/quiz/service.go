package quiz

import (
	"errors"
	"fmt"
	"github.com/gokhantamkoc/quiz-app-api/api"
)

type Service struct {
	Quiz map[string]api.Quiz
}

func NewService() *Service {
	return &Service{
		Quiz: make(map[string]api.Quiz),
	}
}


func (s *Service) CreateQuiz(newQuiz *api.Quiz) (*api.Quiz, error) {
	s.Quiz[newQuiz.Id] = *newQuiz
	return newQuiz, nil
}

func (s *Service) UpdateQuiz() (*api.Quiz, error) {
	return &api.Quiz{
		Id: "1",
		Name: "Flutter Quiz",
		Status: 8,
	}, nil
}

func (s *Service) ListQuiz() []api.Quiz {
	quizArr := []api.Quiz{}
	for _, v := range s.Quiz {
		quizArr = append(quizArr, v)
	}
	return quizArr
}

func (s *Service) GetQuiz(quizId string) (*api.Quiz, error) {
	obj, ok :=  s.Quiz[quizId]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Quiz bulunamadı! (Id: %s)", quizId))
	}
	return &obj, nil
}

//fetch("http://localhost:9001/quiz/list")
//.then(response => {
//	alert(response.json)
//})
//.catch(error => {
//	alert(error)
//});