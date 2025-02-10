package main

import (
	controller "SHOP_PORTAL_BACKEND/CONTROLLER"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	util "SHOP_PORTAL_BACKEND/UTILS"

	"github.com/kataras/iris/v12"
)

func main() {

	helper.Onit()

	app := iris.New()

	// to get server is up
	app.Get("/shop/ping", func(ctx iris.Context) {
		helper.ServerUp(ctx)
	})

	// generate token
	app.Post("/shop/generateToken", func(ctx iris.Context) {
		helper.SetApiName(util.GENERATE_TOKEN, ctx)
		controller.GenerateToken(ctx)
	})

	// api to add new shop owner
	app.Post("shop/addShowOwner", func(ctx iris.Context) {
		helper.SetApiName(util.POST_SHOP_OWNER, ctx)
		controller.PostShopOwner(ctx)
	})

	// api to update shop owner
	app.Put("shop/updateShopOwner", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_SHOP_OWNER, ctx)
		controller.PutShopOwner(ctx)
	})

	// api to get shop owner details
	app.Get("shop/getShopOwner", func(ctx iris.Context) {
		helper.SetApiName(util.GET_SHOP_OWNER, ctx)
		controller.GetShopOwner(ctx)
	})

	// api to get all shop owners
	app.Get("shop/getAllShopOwners", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_SHOP_OWNER, ctx)
		controller.GetAllShopOwner(ctx)
	})

	// Start the server on port 8000
	err := app.Listen(":8000")
	if err != nil {
		app.Logger().Fatal(err)
	}
}
