package helper

import (
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"runtime"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
)

func Onit() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores
	utils.SetCodeMap()                   // set code map
	utils.SetApiHeaders()                // set api headers
	utils.SetValidHeaders()              // set valid headers
	utils.SetApiQParams()
	utils.NewLogger()    // new logger
	database.ConnectDB() // connect to database
}

func ServerUp(ctx iris.Context) {
	ctx.HTML("Backend Server Is UP")
}

func SetApiName(apiName string, ctx iris.Context) {
	shop_id := ctx.URLParam("owner_reg_id")
	logprefix := ("[" + time.Now().Format("2006-01-02 15:04:05") + "] ") + apiName + "_SHOP_ID_" + shop_id + " : "
	ctx.Values().Set("logPrefix", logprefix)
	ctx.Values().Set("apiName", apiName)
	utils.Logger.Info(logprefix + "Request Recieved.")
}

func SetCORS(app *iris.Application) {

	// Create a CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{utils.ACCEPT, utils.CONTENT_TYPE, utils.TOKEN, utils.SKIP_TOKEN, utils.ACCEPT_ENCODING, utils.CATCH_CONTROL},
		AllowCredentials: false,
	})

	// Apply the CORS middleware to all routes
	app.Use(func(ctx iris.Context) {
		c.HandlerFunc(ctx.ResponseWriter(), ctx.Request())
		ctx.Next()
	})
}
