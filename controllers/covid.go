package controllers

import (
	"context"
	"github.com/ShubhamBansal1997/covid-app/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Covid struct {
	ID           string    `json:"id"`
	State        string    `json:"state"`
	StateCases   int       `json:"state_cases"`
	CountryCases int       `json:"country_cases"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var collection *mongo.Collection

func CovidCollection(c *mongo.Database) {
	collection = c.Collection("covid")
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Covid  `json:"data,omitempty"`
}

// FetchData Fetch Data godoc
// @Summary Fetch Covid-19 Stats
// @Description fetch covid-19 state-wise data and store in the db
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /fetch-data [get]
func FetchData(c echo.Context) error {
	data, totalCount, err := utils.GetCovidData()
	resp := Response{}
	if err != nil {
		resp.Status = http.StatusBadRequest
		resp.Message = "Unable to fetch data"
		return c.JSON(http.StatusBadRequest, resp)
	}
	for _, v := range data {
		id, err := getCovidEntry(v.State)
		if err != nil {
			_, err = createCovidEntry(v.State, v.TotalCases, totalCount)
			if err != nil {
				log.Printf("Entry Not Created: %v", v.State)
			}
			continue
		}
		_, err = updateCovidEntry(id.ID, v.TotalCases, totalCount)
		if err != nil {
			log.Printf("Entry Not Updated: %v", v.State)
		}
	}
	resp.Status = http.StatusOK
	resp.Message = "Data fetched successfully"
	return c.JSON(http.StatusOK, resp)
}

// GetData Get Data godoc
// @Summary Get Covid-19 Stats
// @Description Get covid-19 stats as per your location (latitude and longitude)
// @Accept  json
// @Produce  json
// @Param lat query string false "latitude"
// @Param long query string false "longitude"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /get-data [get]
func GetData(c echo.Context) error {
	lat := c.QueryParam("lat")
	long := c.QueryParam("long")
	resp := Response{}
	if lat == "" || long == "" {
		resp.Status = http.StatusInternalServerError
		resp.Message = "Missing query params"
		return c.JSON(http.StatusInternalServerError, resp)
	}
	state, err := utils.GetStateData(lat, long)
	if err != nil {
		resp.Status = http.StatusInternalServerError
		resp.Message = "Unable to fetch location data"
		return c.JSON(http.StatusInternalServerError, resp)
	}
	covidEntry, err := getCovidEntry(state)
	if err != nil {
		resp.Status = http.StatusNotFound
		resp.Message = "Location data not present"
		return c.JSON(http.StatusNotFound, resp)
	}
	resp.Status = http.StatusOK
	resp.Message = "Resource Fetched"
	resp.Data = covidEntry
	return c.JSON(http.StatusOK, resp)
}

func createCovidEntry(state string, totalCases int, totalCount int) (bool, error) {
	id := uuid.NewV4().String()

	newCovidEntry := Covid{
		ID:           id,
		State:        state,
		StateCases:   totalCases,
		CountryCases: totalCount,
		UpdatedAt:    time.Now(),
	}
	_, err := collection.InsertOne(context.TODO(), newCovidEntry)
	if err != nil {
		log.Printf("Error while inserting new covid entry into db, Reason: %v\n", err)
		return false, err
	}
	return true, nil
}

func getCovidEntry(state string) (Covid, error) {
	var covidEntry Covid
	err := collection.FindOne(context.TODO(), bson.M{"state": state}).Decode(&covidEntry)
	if err != nil {
		log.Printf("Error while getting covid entry, Reason: %v\n", err)
		return Covid{}, err
	}
	return covidEntry, nil
}

func updateCovidEntry(id string, totalCases int, totalCount int) (bool, error) {
	newData := bson.M{
		"$set": bson.M{
			"statecases":   totalCases,
			"countrycases": totalCount,
			"updatedat":    time.Now(),
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": id}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		return false, err
	}
	return true, nil
}
