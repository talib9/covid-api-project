package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/config"
	"project/logs"
	// "project/config.json"
)

//Structure of the geolocation data from api.mapmyindia.com

type Document struct {
	ResponseCode int                 `json:"responseCode"`
	Results      []map[string]string `json:"results"`
	Version      string              `json:"version"`
}

// Taking the inputs from the user and fetching the data for those specific coodirnates. It returns the state
func GetByGeoCoordinates(lat string, lng string) (string, error) {

	configuration := config.GetConfigurations()
	url := configuration.Geocoding_api
	api := configuration.Geocoding_apikey
	query := fmt.Sprintf("%s%s/rev_geocode?lat=%s&lng=%s&region=IND", url, api, lat, lng)
    // fmt.Println("query= " ,query)
	
	// query = "https://apis.mapmyindia.com/advancedmaps/v1/325129ace3489e1600a7cf08d4952970/rev_geocode?lat=28.7041&lng=77.1025&region=IND"
	res, err := http.Get(query)
	// fmt.Println(res)
	if err != nil {
		logs.MyLogger(err)
	}
	fmt.Println(err)
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(body)
	defer res.Body.Close()
	var json_data Document
	json.Unmarshal([]byte(body), &json_data)
	// fmt.Println(json_data)
	result := json_data.Results[0]
	state := result["state"]
	// fmt.Println(state)
	return state, err
}
