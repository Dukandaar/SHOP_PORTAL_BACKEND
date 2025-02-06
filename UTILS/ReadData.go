package utils

import "github.com/kataras/iris/v12"

func ReadHeader(ctx iris.Context) map[string]interface{} {
	header := make(map[string]interface{})
	for key, val := range ctx.Request().Header {
		header[key] = val[0]
	}
	return header
}

func ReadQParams(ctx iris.Context) map[string]interface{} {
	qparams := make(map[string]interface{})
	for key, val := range ctx.URLParams() {
		qparams[key] = val
	}
	return qparams
}

func ReadReqBody(ctx iris.Context) map[string]interface{} {
	body := make(map[string]interface{})
	ctx.ReadJSON(&body)
	return body
}
