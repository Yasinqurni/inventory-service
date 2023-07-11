package response

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func NewErrorResponse(message string, err string) *ErrorResponse {
	return &ErrorResponse{
		Error:   true,
		Message: message,
		Err:     err,
	}
}

type ErrorResponseValidate struct {
	Error   bool          `json:"error"`
	Message DescribeError `json:"message"`
	Err     string        `json:"err,omitempty"`
}

type DescribeError struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Quantity  string `json:"quantity"`
	SkuNumber string `json:"sku_number"`
	Notes     string `json:"notes"`
}

func NewErrorResponseValidate(describeError DescribeError, err string) *ErrorResponseValidate {
	return &ErrorResponseValidate{
		Error:   true,
		Message: describeError,
		Err:     err,
	}
}
