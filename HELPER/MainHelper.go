package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"runtime"

	"github.com/kataras/iris/v12"
)

func Onit() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores
	utils.SetCodeMap()                   // set code map
	utils.SetApiHeaders()                // set api headers
	utils.SetValidHeaders()              // set valid headers
	utils.NewLogger()                    // new logger
	database.ConnectDB()                 // connect to database
}

func ServerUp(ctx iris.Context) {
	ctx.HTML("Backend Server Is UP")
}

func SetApiName(apiName string, ctx iris.Context) {
	shop_id := ctx.URLParam("Shop_id")
	logprefix := apiName + "_SHOP_ID_" + shop_id + " : "
	ctx.Values().Set("logPrefix", logprefix)
	ctx.Values().Set("apiName", apiName)
	utils.Logger.Info(logprefix + "Request Recieved.")
}
