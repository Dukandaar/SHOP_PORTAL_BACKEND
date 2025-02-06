package validator

func ValidateHeader(header map[string]interface{}) error {
	return nil
}

func ValidateQParams(qparams map[string]interface{}) error {
	return nil
}

func ValidateBody(body map[string]interface{}) error {
	return nil
}

func ValidateData(header map[string]interface{}, qparmas map[string]interface{}, reqBody map[string]interface{}) (string, error) {

	headerError := ValidateHeader(header)
	if headerError != nil {
		return "INVALID_HEADER", headerError
	}

	qparamsError := ValidateQParams(qparmas)
	if qparamsError != nil {
		return "INVALID_QPARAMS", qparamsError
	}

	bodyError := ValidateBody(reqBody)
	if bodyError != nil {
		return "INVALID_BODY", bodyError
	}

	return "SUCCESS", nil
}
