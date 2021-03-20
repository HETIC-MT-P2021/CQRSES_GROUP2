package response

// Response defines the structure of the response.
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Status defines the structure of the status entities for the response.
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewResponse create the structure of the JSON response.
func NewResponse(code int, message string, data interface{}) *Response {
	res := &Response{
		Data: data,
		Status: Status{
			Code:    code,
			Message: message,
		},
	}

	return res
}
