package docs

import (
	"bytes"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
      "contact": {
        "email": "talibansari96487@gmail.com",
        "name": "Covid Api Support",
        "url": "http://www.covidapi.io/support"
      },
      "description": "This is a covid data api. It fetches data from external apis and gives the data of a particular state.",
      "license": {
        "name": "Apache 2.0",
        "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
      },
      "termsOfService": "http://swagger.io/terms/",
      "title": "Covid Data Api",
      "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
      "/GetStateData": {
        "get": {
          "consumes": [
            "application/json"
          ],
          "description": "Accept state as query parameter and return the data of that particular state. To get Data of India Enter India in state.",
          "parameters": [
            {
              "description": "Enter State",
              "in": "query",
              "name": "state",
              "required": true,
              "type": "string"
            }
          ],
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "OK",
              "headers": {
                "Token": {
                  "description": "qwerty",
                  "type": "string"
                }
              },
              "schema": {
                "$ref": "#/definitions/model.OneState"
              }
            },
            "400": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "404": {
              "description": "Not Found",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "500": {
              "description": "Internal Server Error",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            }
          },
          "summary": "Get data by state",
          "tags": [
            "Api"
          ]
        }
      },
      "/GetAllData": {
        "get": {
          "consumes": [
            "application/json"
          ],
          "description": "This returns data of all states. Date of India is also returned.",
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "OK",
              "headers": {
                "Token": {
                  "description": "qwerty",
                  "type": "string"
                }
              },
              "schema": {
                "$ref": "#/definitions/model.AllState"
              }
            },
            "400": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "404": {
              "description": "Not Found",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "500": {
              "description": "Internal Server Error",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            }
          },
          "summary": "Data of all states and UTs of India.",
          "tags": [
            "Api"
          ]
        }
      },
      "/GetByGeoLocation": {
        "get": {
          "consumes": [
            "application/json"
          ],
          "description": "It accepts the geolocation in query parameter. If the geolocation is of India it will return the data of the state corresponding to the geolocation entered",
          "parameters": [
            {
              "in": "query",
              "name": "latitude",
              "description": " Pass a latitude in query",
              "required": true,
              "type": "string"
            },
            {
              "in": "query",
              "name": "longitude",
              "description": " Pass a longitude in query",
              "required": true,
              "type": "string"
            }
          ],
          "produces": [
            "application/json"
          ],
          "responses": {
            "200": {
              "description": "OK",
              "headers": {
                "Token": {
                  "description": "qwerty",
                  "type": "string"
                }
              },
              "schema": {
                "$ref": "#/definitions/model.OneState"
              }
            },
            "400": {
              "description": "Bad Request",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "404": {
              "description": "Not Found",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            },
            "500": {
              "description": "Internal Server Error",
              "schema": {
                "$ref": "#/definitions/httputil.HTTPError"
              }
            }
          },
          "summary": "Data of state based on geolocation",
          "tags": [
            "Api"
          ]
        }
      }
    },
    "definitions": {
        "httputil.HTTPError": {
          "properties": {
            "code": {
              "example": 400,
              "type": "integer"
            },
            "message": {
              "example": "status bad request",
              "type": "string"
            }
          },
          "type": "object"
        },
        "model.OneState": {
          "type": "object",
          "required": [
            "state_code",
            "total_cases",
            "total_recovered",
            "total_death",
            "total_vaccinated1",
            "total_vaccinated2",
            "total_tested",
            "last_updated"
          ],
          "properties": {
            "state_code": {
              "type": "string",
              "example": "HR"
            },
            "total_cases": {
              "type": "integer",
              "example": 123456
            },
            "total_recovered": {
              "type": "integer",
              "example": 120356
            },
            "total_death": {
              "type": "integer",
              "example": 3654
            },
            "total_vaccinated1": {
              "type": "integer",
              "example": 658545
            },
            "total_vaccinated2": {
              "type": "integer",
              "example": 365244
            },
            "total_tested": {
              "type": "integer",
              "example": 2356878
            },
            "last_updated": {
              "type": "string",
              "format": "date-time",
              "example": "2016-08-29T09:12:33.001Z"
            }
          }
        },
        "model.AllState": {
          "type": "object",
          "required": [
            "state_code",
            "total_cases",
            "total_recovered",
            "total_death",
            "total_vaccinated1",
            "total_vaccinated2",
            "total_tested",
            "last_updated"
          ],
          "properties": {
            "state_code": {
              "type": "string",
              "example": "HR"
            },
            "total_cases": {
              "type": "integer",
              "example": 123456
            },
            "total_recovered": {
              "type": "integer",
              "example": 120356
            },
            "total_death": {
              "type": "integer",
              "example": 3654
            },
            "total_vaccinated1": {
              "type": "integer",
              "example": 658545
            },
            "total_vaccinated2": {
              "type": "integer",
              "example": 365244
            },
            "total_tested": {
              "type": "integer",
              "example": 2356878
            },
            "last_updated": {
              "type": "string",
              "format": "date-time",
              "example": "2016-08-29T09:12:33.001Z"
            }
          }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
