package main

import (
	helper "SHOP_PORTAL_BACKEND/HELPER"

	"github.com/kataras/iris/v12"
)

func main() {

	helper.Onit()

	app := iris.New()

	// to get server is up
	app.Get("/shop/ping", func(ctx iris.Context) {
		helper.ServerUp(ctx)
	})

	// Start the server on port 8000
	err := app.Listen(":8000")
	if err != nil {
		app.Logger().Fatal(err)
	}
}
