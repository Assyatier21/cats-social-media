package models

type SuccessSwaggerResp struct {
	Code    int         `json:"code" example:"200"`
	Status  string      `json:"status" example:"OK"`
	Message string      `json:"message" example:"Succeed"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type BadRequestSwaggerResp struct {
	Code    int         `json:"code" example:"400"`
	Status  string      `json:"status" example:"Bad Request"`
	Message string      `json:"message" example:"Wrong or Unexpected Request"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors" example:"{field} is required"`
}

type InternalServerErrorSwaggerResp struct {
	Code    int         `json:"code" example:"500"`
	Status  string      `json:"status" example:"Internal Server Error"`
	Message string      `json:"message" example:"Unable to Handle This Request"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type NotFoundSwaggerResp struct {
	Code    int         `json:"code" example:"404"`
	Status  string      `json:"status" example:"Not Found"`
	Message string      `json:"message" example:"Request Not Found"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
