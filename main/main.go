package main

import (
	"fmt"
	"project/app/api"
	"project/app/handlers"
	"project/mongodb"

	_ "project/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := echo.New()
	// fmt.Println("ign")
	final_data := api.GettingData()
	// fmt.Println(final_data)

	mongodb.UpdatingData(final_data)

	e.GET("GetStateData", handlers.GetCases)
	e.GET("/GetAllData", handlers.GetAllCases)
	e.GET("/GetByGeoLocation", handlers.GetDataFromGeoLocation)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// port := os.Getenv("PORT")

	address := fmt.Sprintf("%s:%s", "0.0.0.0", "8080")
	e.Logger.Fatal(e.Start(address))
}
