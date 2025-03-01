package helper

import (
	config "SHOP_PORTAL_BACKEND/CONFIG"
	database "SHOP_PORTAL_BACKEND/DATABASE"
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"runtime"

	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
)

func Onit() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores
	utils.SetCodeMap()                   // set code map
	utils.SetApiHeaders()                // set api headers
	utils.SetValidHeaders()              // set valid headers
	utils.SetApiQParams()
	utils.NewLogger()                    // new logger
	config.ReadAllEnvironmentVariables() // read all environment variables
	database.ConnectDB()                 // connect to database
}

func ServerUp(ctx iris.Context) {
	ctx.HTML("Backend Server Is UP")
}

func SetApiName(apiName string, ctx iris.Context) {
	var logPrefix string
	shop_id := ctx.URLParam("owner_reg_id")
	if shop_id != utils.NULL_STRING {
		logPrefix = apiName + " : " + shop_id + " : "
	} else {
		logPrefix = apiName + " : "
	}

	ctx.Values().Set("logPrefix", logPrefix)
	ctx.Values().Set("apiName", apiName)
	utils.Logger.Info(logPrefix + "Request Recieved.")
}

func SetCORS(app *iris.Application) {

	// Create a CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{utils.ACCEPT, utils.CONTENT_TYPE, utils.TOKEN, utils.SKIP_TOKEN, utils.ACCEPT_ENCODING, utils.CACHE_CONTROL},
		AllowCredentials: false,
	})

	// Apply the CORS middleware to all routes
	app.Use(func(ctx iris.Context) {
		c.HandlerFunc(ctx.ResponseWriter(), ctx.Request())
		ctx.Next()
	})
}
