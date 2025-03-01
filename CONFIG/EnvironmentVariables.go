package config

import (
	utils "SHOP_PORTAL_BACKEND/UTILS"
	"os"

	"github.com/joho/godotenv"
)

var DRIVER_NAME string
var DB_USER string
var DB_PASSWORD string
var DB_HOST string
var DB_PORT string
var DB_NAME string
var PRIVATE_KEY string
var PUBLIC_KEY string
var JWT_SECRET string
var DISABLE_COLORS string

func ReadAllEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	DRIVER_NAME = os.Getenv("DRIVER_NAME")
	if DRIVER_NAME == utils.NULL_STRING {
		utils.Logger.Warn("DRIVER_NAME is not set")
	}

	DB_USER = os.Getenv("DB_USER")
	if DB_USER == utils.NULL_STRING {
		utils.Logger.Warn("DB_USER is not set")
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == utils.NULL_STRING {
		utils.Logger.Warn("DB_PASSWORD is not set")
	}

	DB_HOST = os.Getenv("DB_HOST")
	if DB_HOST == utils.NULL_STRING {
		utils.Logger.Warn("DB_HOST is not set")
	}

	DB_PORT = os.Getenv("DB_PORT")
	if DB_PORT == utils.NULL_STRING {
		utils.Logger.Warn("DB_PORT is not set")
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == utils.NULL_STRING {
		utils.Logger.Warn("DB_NAME is not set")
	}

	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	if PRIVATE_KEY == utils.NULL_STRING {
		utils.Logger.Warn("PRIVATE_KEY is not set")
	}

	PUBLIC_KEY = os.Getenv("PUBLIC_KEY")
	if PUBLIC_KEY == utils.NULL_STRING {
		utils.Logger.Warn("PUBLIC_KEY is not set")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == utils.NULL_STRING {
		utils.Logger.Warn("JWT_SECRET is not set")
	}

	DISABLE_COLORS = os.Getenv("DISABLE_COLORS")
	if DISABLE_COLORS == utils.NULL_STRING {
		utils.Logger.Warn("DISABLE_COLORS is not set")
	}

	utils.Logger.Info("Environment variables read successfully")
}
