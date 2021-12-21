package config

import (
	"context"
	"project/logs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
	
// Structure for configuartion from github.config file
type Configurations struct {
	Mongodb_user     string `json:"mongodb_user" bson:"mongodb_user"`
	Mongodb_password string `json:"mongodb_password" bson:"mongodb_password"`
	Mongodb_host     string `json:"mongodb_host" bson:"mongodb_host"`
	Geocoding_apikey string `json:"geocoding_apikey" bson:"geocoding_apikey"`
	Covid_api        string `json:"covid_api" bson:"covid_api"`
	Geocoding_api    string `json:"geocoding_api" bson:"geocoding_api"`
}

// Structure for configuration for covid data to be stored in mongodb
type ResponseData struct {
	// ID                primitive.ObjectID `json:"_id" bson:"_id" `
	State_code        string    `json:"state_code" bson:"state_code"`
	Total_cases       int       `json:"total_cases" bson:"total_cases"`
	Total_recovered   int       `json:"total_recovered" bson:"total_recovered"`
	Total_death       int       `json:"total_death" bson:"total_death"`
	Total_vaccinated1 int       `json:"total_vaccinated1" bson:"total_vaccinated1"`
	Total_vaccinated2 int       `json:"total_vaccinated2" bson:"total_vaccinated2"`
	Total_tested      int       `json:"total_tested" bson:"total_tested"`
	Last_updated      time.Time `json:"last_updated" bson:"last_updated"`
}

// States code
var state_code map[string]string = map[string]string{"Total": "TT",
	"Andaman and Nicobar": "AN",
	"Andhra Pradesh":      "AP",
	"Arunachal Pradesh":   "AR",
	"Assam":               "AS",
	"Bihar":               "BR",
	"Chandigarh":          "CH",
	"Chhattisgarh":        "CT",
	"Delhi":               "DL",
	"Goa":                 "GA",
	"Gujarat":             "GJ",
	"Haryana":             "HR",
	"Himachal Pradesh":    "HP",
	"Jammu and Kashmir":   "JK",
	"Jharkhand":           "JH",
	"Karnataka":           "KA",
	"Kerala":              "KL",
	"Ladakh":              "LA",
	"Lakshadweep":         "LD",
	"Madhya Pradesh":      "MP",
	"Maharashtra":         "MH",
	"Manipur":             "MN",
	"Meghalaya":           "ML",
	"Mizoram":             "MZ",
	"Nagaland":            "NL",
	"Odisha":              "OR",
	"Puducherry":          "PY",
	"Punjab":              "PB",
	"Rajasthan":           "RJ",
	"Sikkim":              "SK",
	"Tamil Nadu":          "TN",
	"Telangana":           "TG",
	"Tripura":             "TR",
	"Uttar Pradesh":       "UP",
	"Uttarakhand":         "UT",
	"West Bengal":         "WB"}

// Get State Codes
func GetStateCodes() map[string]string {
	return state_code
}

// Get State Name from State Code
func StateCodeFromStateName() map[string]string {
	state_name := make(map[string]string)
	for k, v := range state_code {
		state_name[v] = k
	}
	return state_name
}

// Connecting to mongoDb
func ConnectionMongoDb() (*mongo.Collection, error) {
	configurations := GetConfigurations()
	mongodb_user := configurations.Mongodb_user
	mongodb_password := configurations.Mongodb_password
	mongodb_Host := configurations.Mongodb_host
	url := fmt.Sprintf("mongodb+srv://%s:%s@%s/test?retryWrites=true&w=majority", mongodb_user, mongodb_password, mongodb_Host)
	
	// url := fmt.Sprintf("mongodb+srv://Talib:Talib96487@cluster0.cvwln.mongodb.net/test?retryWrites=true&w=majority")
	// fmt.Println(url)
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logs.MyLogger(err)
		log.Panic(err)
	}
	db := client.Database("covid")
	collection := db.Collection("statewise")
    // fmt.Println(collection)
	return collection, nil
}

// Getting configurations
func GetConfigurations() Configurations {

	config, err := http.Get("https://raw.githubusercontent.com/talib9/covid-api-project/master/config.json")
	if err != nil {
		logs.MyLogger(err)
		log.Panic(err)
	}
	body, _ := ioutil.ReadAll(config.Body)
	defer config.Body.Close()
	var configuration Configurations
	json.Unmarshal([]byte(body), &configuration)
	return configuration
}