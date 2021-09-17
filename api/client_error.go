package api

// ValidationError => Kullanici
// UnexpectedError => FrontEnd dev

type ClientError struct {
	Message string `json:"message"`
}