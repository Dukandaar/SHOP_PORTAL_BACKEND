package structs

type ErrorResponse struct {
	Stat             string           `json:"stat"`
	ErrorSubResponse ErrorSubResponse `json:"rsp"`
}

type ErrorSubResponse struct {
	ErrorCode       int    `json:"code"`
	ErrorMsg        string `json:"msg"`
	ErrorDescrition string `json:"description"`
}

type SuccessResponse struct {
	Stat               string             `json:"stat"`
	SuccessSubResponse SuccessSubResponse `json:"rsp"`
}

type SuccessSubResponse struct {
	SuccessMsg string `json:"msg"`
}

type TokenResponse struct {
	Stat             string           `json:"stat"`
	TokenSubResponse TokenSubResponse `json:"rsp"`
}

type TokenSubResponse struct {
	Token string `json:"token"`
}
