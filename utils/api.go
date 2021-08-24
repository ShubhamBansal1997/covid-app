package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type CovidStats struct {
	Delta struct {
		Confirmed int `json:"confirmed"`
	} `json:"delta"`
	Delta2114 struct {
		Confirmed int `json:"confirmed"`
	} `json:"delta21_14"`
	Delta7 struct {
		Confirmed   int `json:"confirmed"`
		Recovered   int `json:"recovered"`
		Tested      int `json:"tested"`
		Vaccinated1 int `json:"vaccinated1"`
		Vaccinated2 int `json:"vaccinated2"`
	} `json:"delta7"`
	Districts struct {
		Nicobars struct {
			Delta7 struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"delta7"`
			Meta struct {
				Population int `json:"population"`
			} `json:"meta"`
			Total struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"total"`
		} `json:"Nicobars"`
		NorthAndMiddleAndaman struct {
			Delta7 struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"delta7"`
			Meta struct {
				Population int `json:"population"`
			} `json:"meta"`
			Total struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"total"`
		} `json:"North and Middle Andaman"`
		SouthAndaman struct {
			Delta7 struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"delta7"`
			Meta struct {
				Population int `json:"population"`
			} `json:"meta"`
			Total struct {
				Vaccinated1 int `json:"vaccinated1"`
				Vaccinated2 int `json:"vaccinated2"`
			} `json:"total"`
		} `json:"South Andaman"`
		Unknown struct {
			Delta struct {
				Confirmed int `json:"confirmed"`
			} `json:"delta"`
			Delta2114 struct {
				Confirmed int `json:"confirmed"`
			} `json:"delta21_14"`
			Delta7 struct {
				Confirmed int `json:"confirmed"`
				Recovered int `json:"recovered"`
			} `json:"delta7"`
			Total struct {
				Confirmed int `json:"confirmed"`
				Deceased  int `json:"deceased"`
				Recovered int `json:"recovered"`
			} `json:"total"`
		} `json:"Unknown"`
	} `json:"districts"`
	Meta struct {
		Date        string    `json:"date"`
		LastUpdated time.Time `json:"last_updated"`
		Population  int       `json:"population"`
		Tested      struct {
			Date   string `json:"date"`
			Source string `json:"source"`
		} `json:"tested"`
	} `json:"meta"`
	Total struct {
		Confirmed   int `json:"confirmed"`
		Deceased    int `json:"deceased"`
		Recovered   int `json:"recovered"`
		Tested      int `json:"tested"`
		Vaccinated1 int `json:"vaccinated1"`
		Vaccinated2 int `json:"vaccinated2"`
	} `json:"total"`
}

type CovidData struct {
	State string
	TotalCases int
}

type LocationApiData struct {
	Data []struct {
		Latitude           float64     `json:"latitude"`
		Longitude          float64     `json:"longitude"`
		Type               string      `json:"type"`
		Distance           float64     `json:"distance"`
		Name               string      `json:"name"`
		Number             interface{} `json:"number"`
		PostalCode         interface{} `json:"postal_code"`
		Street             interface{} `json:"street"`
		Confidence         float64     `json:"confidence"`
		Region             string      `json:"region"`
		RegionCode         string      `json:"region_code"`
		County             string      `json:"county"`
		Locality           string      `json:"locality"`
		AdministrativeArea interface{} `json:"administrative_area"`
		Neighbourhood      interface{} `json:"neighbourhood"`
		Country            string      `json:"country"`
		CountryCode        string      `json:"country_code"`
		Continent          string      `json:"continent"`
		Label              string      `json:"label"`
	} `json:"data"`
}

func GetCovidData() ([]CovidData, int, error) {
	baseUrl := os.Getenv("COVID_API")
	response, err := http.Get(baseUrl)
	if err != nil {
		return nil, 0, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	var covidApiData map[string]CovidStats
	err = json.Unmarshal(responseData, &covidApiData)
	if err != nil {
		return nil, 0, err
	}
	var covidData []CovidData
	totalCount := 0
	for k, v := range covidApiData {
		state, err := GetStateName(k)
		totalCount = totalCount + v.Total.Confirmed
		if err != nil {
			log.Printf("Error while getting State Name%v\n", err)
			continue
		}
		newData := CovidData{
			State: state,
			TotalCases: v.Total.Confirmed,
		}
		covidData = append(covidData, newData)
	}
	return covidData, totalCount, nil
}

func GetStateData(lat string, long string) (string, error) {
	log.Print("Came here")
	baseUrl := os.Getenv("GEOLOCATION_API")
	api := baseUrl + lat + "," + long
	response, err := http.Get(api)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var locationApiData LocationApiData
	err = json.Unmarshal(responseData, &locationApiData)
	if err != nil {
		return "", err
	}
	if len(locationApiData.Data) > 0 {
		return locationApiData.Data[0].Region, nil
	}
	return "", errors.New("no location found")
}
