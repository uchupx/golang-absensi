package dto

type ErrorResponse struct {
	HTTPCode int32  `json:"-"`
	Code     int32  `json:"code"`
	Message  string `json:"message"`
}

type ErrorMapping struct {
	Status        int
	ErrorResponse ErrorResponse
}

type BaseResponse struct {
	Data  interface{}    `json:"data"`
	Error *ErrorResponse `json:"error"`
}

type ResponseParam struct {
	Status  int
	Payload BaseResponse
}

type BaseListResponse struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}
