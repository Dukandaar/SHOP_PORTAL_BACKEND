package utils

type Codes struct {
	StatusCode  int
	Message     string
	Description string
}

const (
	StatusOK                  = 200
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusRequestTimeout      = 408
	StatusInternalServerError = 500
)

var CodeMap map[string]Codes

func SetCodeMap() {
	CodeMap = make(map[string]Codes)
	CodeMap["200001"] = Codes{200001, "OK", "Valid Request"}
	CodeMap["400001"] = Codes{400001, "Invalid headers", "Invalid headers"}
	CodeMap["400002"] = Codes{400002, "Invalid Qparams", "Invalid Query Parameters"}
	CodeMap["400003"] = Codes{400003, "Invalid Body", "Malformed ReqBody Content"}
	CodeMap["404001"] = Codes{404001, "4002", "Not Found"}
	CodeMap["408001"] = Codes{408001, "4003", "Request Timeout"}
	CodeMap["500001"] = Codes{500001, "Internal Server Error", "Internal Server Error"}
}
