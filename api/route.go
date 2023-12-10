package api

import "github.com/gin-gonic/gin"

type Config struct{
	Router *gi.Engine
}

func (app *Config) Routes(){
	app.Routes.POST("/otp", app.sendSMS())
	app.Routes.POST("/verifyOTP", app.verifySMS())
}