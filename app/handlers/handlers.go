package handlers

import (
	"project/app/api"
	"project/config"
	"project/mongodb"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// All the handlers for the routes

// Get Cases ouputs for one specific state based on the state query passed in url
// endpoint is /GetCases
func GetCases(c echo.Context) error {
	valid := c.Request().URL.Query().Has("state")

	if !valid {
		return c.JSON(http.StatusBadRequest, "Invalid query string key name it should be state")
	}
	state := c.QueryParam("state")
	state_code, present := config.GetStateCodes()[state]
	if !present {
		states := config.GetStateCodes()
		state_list := ""
		for k := range states {
			state_list = state_list + k + ","
		}
		list := fmt.Sprint("Please provide a valid state name... Here is the list of availabe state name........", state_list)
		return c.JSON(http.StatusBadRequest, list)
	}
	data, err := mongodb.GetData(state_code)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, data)
}

// Get Cases ouputs for all the  states in the mongoDb
// endpoint is /GetAllCases
func GetAllCases(c echo.Context) error {
	data, err := mongodb.GetAllData()
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)
}

// Get Cases ouputs for one specific state based on the Geolocation passed in url
// endpoint is /GetDataFromGeoLocation

func GetDataFromGeoLocation(c echo.Context) error {
	latitude := c.Request().URL.Query().Has("latitude")
	longitude := c.Request().URL.Query().Has("longitude")


	if !latitude || !longitude {
		
		return c.JSON(http.StatusBadRequest, "Invalid query string key name it should be lat and lng")
	}
	lat := c.QueryParam("latitude")
	lng := c.QueryParam("longitude")
   
	state, err := api.GetByGeoCoordinates(lat, lng)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Please give coordinates of India")
	}
	data, err := mongodb.GetData(config.GetStateCodes()[state])
	if err != nil {
		return c.JSON(http.StatusNotFound, "Invaild StateName")
	}
	return c.JSON(http.StatusOK, data)

}