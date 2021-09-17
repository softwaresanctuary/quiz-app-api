package api

type Quiz struct {
	Id string 			`json:"id" binding:"required"`
	Name string 		`json:"name" binding:"required"`
	Status uint8 		`json:"status" binding:"required"`
	Category string 	`json:"category,omitempty"`
}

type QuizSearch struct {
	Field	string `json:"field" binding:"required"`
	Value	string `json:"value" binding:"required"`
}