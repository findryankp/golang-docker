package main

import (
	"be15/client/planets"
)

func main() {
	// response, err := http.Get("https://swapi.dev/api/planets/1")
	// if err != nil {
	// 	log.Fatal("error consume swapi api")
	// }

	// responseBody, _ := ioutil.ReadAll(response.Body)
	// defer response.Body.Close()

	// var data planets.Planet
	// errJson := json.Unmarshal(responseBody, &data)
	// if errJson != nil {
	// 	log.Fatal("error unmarshal")
	// }

	// fmt.Println(data.Name)
	// fmt.Println(data.RotationPeriod)
	// fmt.Println(data.OrbitalPeriod)
	// fmt.Println(data.Residents)
	// for _, value := range data.Residents {
	// 	fmt.Println("people: ", value)
	// }

	planets.GetAllPlanets()
}
