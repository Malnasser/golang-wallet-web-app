package core

type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request data"`
	Message string `json:"details,omitempty" example:"Field validation failed"`
} // @name ErrorResponse

type PaginationRequest struct {
	Page     int `json:"page" form:"page" example:"1"`
	PageSize int `json:"page_size" form:"page_size" example:"10"`
} // @name PaginationRequest

type PaginationResponse struct {
	Page       int   `json:"page" example:"1"`
	PageSize   int   `json:"page_size" example:"10"`
	TotalCount int64 `json:"total_count" example:"100"`
	TotalPages int   `json:"total_pages" example:"10"`
} // @name PaginationResponse
