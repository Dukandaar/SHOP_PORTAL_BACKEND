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
	CodeMap["400001"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [MISSING HEADER]"}
	CodeMap["400002"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [INVLAID HEADER]"}
	CodeMap["400003"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [MISSING QPARAM]"}
	CodeMap["400004"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [INVLAID QPARAM]"}
	CodeMap["400005"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [MISSING BODY FIELD]"}
	CodeMap["400006"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [INVLAID BODY FIELD]"}
	CodeMap["400007"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [INVALID BODY FIELD LENGTH]"}
	CodeMap["400008"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [MALFORMED BODY CONTENT]"}
	CodeMap["400009"] = Codes{StatusCode: StatusBadRequest, Message: "Bad Request [ALREADY EXIST]"}
	CodeMap["404001"] = Codes{StatusCode: StatusNotFound, Message: "Data Not Found"}
	CodeMap["500001"] = Codes{StatusCode: StatusInternalServerError, Message: "Internal Server Error"}
}
