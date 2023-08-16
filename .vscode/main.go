package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	_ "embed"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var ginLambda *ginadapter.GinLambda

//go:embed prefectures.json
var prefecturesJSON []byte

//go:embed cities.json
var citiesJSON []byte

//go:embed prefectures-stations.json
var prefectureStationsJSON []byte

type City struct {
	ID                string `json:"id"`
	PrefectureID      int    `json:"prefecture_id"`
	CityEN            string `json:"city_en"`
	CityJA            string `json:"city_ja"`
	SpecialDistrictJA string `json:"special_district_ja"`
}

type PrefectureTrainStation struct {
	ID    int               `json:"id"`
	Name  map[string]string `json:"name"`
	Lines []Line            `json:"lines"`
}

type Prefectures struct {
	ID     int    `json:"id"`
	Kanji  string `json:"prefecture_kanji"`
	Romaji string `json:"prefecture_romaji"`
	Kana   string `json:"prefecture_kana"`
}

type Line struct {
	ID       int               `json:"id"`
	Name     map[string]string `json:"name"`
	Lat      float64           `json:"lat"`
	Lng      float64           `json:"lng"`
	Zoom     int               `json:"zoom"`
	Stations []interface{}     `json:"stations"`
}

func readJSONFile(filename string, data interface{}) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		return err
	}

	return nil
}

func filterPrefectureByID(prefID string, data []PrefectureTrainStation) PrefectureTrainStation {
	var prefecture PrefectureTrainStation

	for _, pref := range data {
		if strconv.Itoa(pref.ID) == prefID {
			prefecture = pref
			break
		}
	}

	return prefecture
}

func getCitiesByPrefectureID(prefectureID string, cities []City) []City {
	var filteredCities []City

	for _, city := range cities {
		if strconv.Itoa(city.PrefectureID) == prefectureID {
			filteredCities = append(filteredCities, city)
		}
	}

	return filteredCities
}

func loadData(data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		log.Fatalf("Error unmarshaling JSON: %s", err)
	}
}

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	// Apply CORS middleware with custom configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Load prefectures data
	var prefecturesData []Prefectures
	loadData(prefecturesJSON, &prefecturesData)

	// Load cities data
	var citiesData []City
	loadData(citiesJSON, &citiesData)

	// Load train stations data
	var prefectureTrainStationsData []PrefectureTrainStation
	loadData(prefectureStationsJSON, &prefectureTrainStationsData)


	// Endpoint to get all Japan prefectures and their cities
	r.GET("/cities-by-prefecture-id", func(c *gin.Context) {
		prefectureID := c.Query("prefecture_id")

		// Check if the prefecture_id is provided and valid
		if prefectureID == "" {
			c.JSON(200, gin.H{
				"error": "prefecture_id is required",
			})
			return
		}

		// Get cities for the specified prefecture_id
		filteredCities := getCitiesByPrefectureID(prefectureID, citiesData)

		// Return the filtered cities
		c.JSON(200, gin.H{
			"cities": filteredCities,
		})
	})

	// Endpoint to get all Japan prefectures and their cities
	r.GET("/prefectures", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"prefectures": prefecturesData,
		})
	})

	// Endpoint to get all Japan prefectures and their train stations
	r.GET("/train-stations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"trainStations": prefectureTrainStationsData,
		})
	})

	// api/lines-by-pref?prefecture_id=1
	r.GET("/lines-by-pref", func(c *gin.Context) {
		prefID := c.Query("prefecture_id")

		if prefID == "" {
			c.JSON(200, gin.H{
				"error": "prefecture_id is required",
			})
			return
		}

		linesForPrefecture := filterPrefectureByID(prefID, prefectureTrainStationsData)
		c.JSON(200, gin.H{
			"prefecture_lines": linesForPrefecture,
		})
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
