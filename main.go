package main

import (
	controller "SHOP_PORTAL_BACKEND/CONTROLLER"
	helper "SHOP_PORTAL_BACKEND/HELPER"
	util "SHOP_PORTAL_BACKEND/UTILS"
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.Default()

	helper.Onit()
	helper.SetCORS(app)

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
	app.Post("shop/addShopOwner", func(ctx iris.Context) {
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

	// get all shop owners
	app.Get("shop/getAllShopOwners", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_SHOP_OWNER, ctx)
		controller.GetAllShopOwner(ctx)
	})

	// add Customer
	app.Post("shop/addCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.POST_CUSTOMER, ctx)
		controller.PostCustomer(ctx)
	})

	// update Customer
	app.Put("shop/updateCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_CUSTOMER, ctx)
		controller.PutCustomer(ctx)
	})

	// get Customer details
	app.Get("shop/getCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.GET_CUSTOMER, ctx)
		controller.GetCustomer(ctx)
	})

	// get all Customer details
	app.Get("shop/getAllCustomers", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_CUSTOMER, ctx)
		controller.GetAllCustomer(ctx)
	})

	// get filtered customer details
	app.Get("shop/getFilteredCustomer", func(ctx iris.Context) {
		helper.SetApiName(util.GET_FILTERED_CUSTOMER, ctx)
		controller.GetFilteredCustomer(ctx)
	})

	// add stock
	app.Post("shop/addStock", func(ctx iris.Context) {
		helper.SetApiName(util.POST_STOCK, ctx)
		controller.PostStock(ctx)
	})

	// update stock
	app.Put("shop/updateStock", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_STOCK, ctx)
		controller.PutStock(ctx)
	})

	// get stock details
	app.Get("shop/getStock", func(ctx iris.Context) {
		helper.SetApiName(util.GET_STOCK, ctx)
		controller.GetStock(ctx)
	})

	// get all stock details
	app.Get("shop/getAllStock", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_STOCK, ctx)
		controller.GetAllStock(ctx)
	})

	// get previous balance
	app.Get("shop/getPreviousBalance", func(ctx iris.Context) {
		helper.SetApiName(util.GET_PREVIOUS_BALANCE, ctx)
		controller.GetPreviousBalance(ctx)
	})

	// api to add customer bill
	app.Post("shop/addCustomerBill", func(ctx iris.Context) {
		helper.SetApiName(util.POST_CUSTOMER_BILL, ctx)
		controller.PostCustomerBill(ctx)
	})

	// api to update customer transaction
	app.Put("shop/updateCustomerTransaction", func(ctx iris.Context) {
		helper.SetApiName(util.PUT_CUSTOMER_TRANSACTION, ctx)
		controller.PutCustomerBill(ctx)
	})

	// api to get Stock history
	app.Get("shop/getStockHistory", func(ctx iris.Context) {
		helper.SetApiName(util.GET_STOCK_HISTORY, ctx)
		controller.GetStockHistory(ctx)
	})

	// // api to get customer bill
	app.Get("shop/getCustomerBill", func(ctx iris.Context) {
		helper.SetApiName(util.GET_CUSTOMER_BILL, ctx)
		controller.GetCustomerBill(ctx)
	})

	// api to get all customer bill
	app.Get("shop/getAllCustomerBill", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_CUSTOMER_BILL, ctx)
		controller.GetAllCustomerBill(ctx)
	})

	// // api to get filtered customer transaction
	// app.Get("shop/getFilteredCustomerTransaction", func(ctx iris.Context) {
	// 	helper.SetApiName(util.GET_FILTERED_CUSTOMER_TRANSACTION, ctx)
	// 	controller.GetFilteredCustomerTransaction(ctx)
	// })

	// api to get all owner bill
	app.Get("shop/getAllOwnerBill", func(ctx iris.Context) {
		helper.SetApiName(util.GET_ALL_OWNER_BILL, ctx)
		controller.GetAllOwnerBill(ctx)
	})

	// api to get previous bill no
	app.Get("shop/getPreviousBillNo", func(ctx iris.Context) {
		helper.SetApiName(util.GET_PREVIOUS_BILL_NO, ctx)
		controller.GetPreviousBillNo(ctx)
	})

	// api to delete stock
	app.Delete("shop/deleteStock", func(ctx iris.Context) {
		helper.SetApiName(util.DELETE_STOCK, ctx)
		controller.DeleteStock(ctx)
	})

	// Start the server on port 8000
	// err := app.Listen(":8000")
	// if err != nil {
	// 	app.Logger().Fatal(err)
	// }

	port := "8000"
	addr := "0.0.0.0:" + port // Bind to 0.0.0.0
	fmt.Println("Server listening on:", addr)

	err := app.Listen(addr) // Start the Iris server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
