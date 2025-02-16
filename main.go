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
		helper.SetApiName(util.SERVER_UP, ctx)
		helper.ServerUp(ctx)
	})

	// generate token
	app.Post("/shop/generateToken", func(ctx iris.Context) {
		helper.SetApiName(util.GENERATE_TOKEN, ctx)
		controller.GenerateToken(ctx)
	})

	// add new shop owner
	app.Post("shop/addShowOwner", func(ctx iris.Context) {
		helper.SetApiName(util.POST_SHOP_OWNER, ctx)
		controller.PostShopOwner(ctx)
	})

	// update shop owner
	app.Put("shop/updateShopOwner", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_SHOP_OWNER, ctx)
		controller.PutShopOwner(ctx)
	})

	// get shop owner details
	app.Get("shop/getShopOwner", func(ctx iris.Context) {
		helper.SetApiName(util.GET_SHOP_OWNER, ctx)
		controller.GetShopOwner(ctx)
	})

	// api to get all shop owners
	app.Get("shop/getAllShopOwners", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_SHOP_OWNER, ctx)
		controller.GetAllShopOwner(ctx)
	})

	// api to add Customer
	app.Post("shop/addCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.POST_CUSTOMER, ctx)
		controller.PostCustomer(ctx)
	})

	// api to update Customer
	app.Put("shop/updateCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_CUSTOMER, ctx)
		controller.PutCustomer(ctx)
	})

	// api to get Customer details
	app.Get("shop/getCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.GET_CUSTOMER, ctx)
		controller.GetCustomer(ctx)
	})

	// api to get all Customers
	app.Get("shop/getAllCustomers", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_CUSTOMER, ctx)
		controller.GetAllCustomer(ctx)
	})

	// api tp get filtered customer
	app.Get("shop/getFilteredCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.GET_FILTERED_CUSTOMER, ctx)
		controller.GetFilteredCustomer(ctx)
	})

	// api to add customer transaction
	app.Post("shop/addCustomerTransaction", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_CUSTOMER_TRANSACTION, ctx)
		controller.PutCustomerTransaction(ctx)
	})

	// Start the server on port 8000
	err := app.Listen(":8000")
	if err != nil {
		app.Logger().Fatal(err)
	}
}
