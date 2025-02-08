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
	CodeMap["200001"] = Codes{StatusCode: StatusOK, Message: "OK"}
	CodeMap["400001"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [HEADER]"}
	CodeMap["400002"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [QUERY-PARAM]"}
	CodeMap["400003"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [BODY]"}
	CodeMap["500001"] = Codes{StatusCode: StatusInternalServerError, Message: "Internal Server Error"}
}
