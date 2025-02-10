package controller

import "github.com/kataras/iris/v12"

func GetCustomer(ctx iris.Context) {

	ctx.JSON("GetCustomer")
}
