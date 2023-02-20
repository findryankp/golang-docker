package planets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PeoplesResponse struct {
	Results []Planet `json:"results"`
	Count   int      `json:"count"`
	Next    string   `json:"next"`
}

type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Residents      []string `json:"residents"`
}

func GetAllPlanets() {
	response, err := http.Get("https://swapi.dev/api/planets")
	if err != nil {
		log.Fatal("error consume swapi api")
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var data PeoplesResponse
	errJson := json.Unmarshal(responseBody, &data)
	if errJson != nil {
		log.Fatal("error unmarshal")
	}

	fmt.Println(data.Results[1].Name)
	for _, value := range data.Results {
		fmt.Println("nama planet:", value.Name)
	}

}
