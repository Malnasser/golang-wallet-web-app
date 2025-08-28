package core

type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request data"`
	Message string `json:"details,omitempty" example:"Field validation failed"`
} // @name ErrorResponse
