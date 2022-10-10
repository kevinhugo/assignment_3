package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/claudiu/gocron"
)

func main() {
	fmt.Println("===================Project Starting===================")
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(UpdateWeatherData)
	<-s.Start()
	fmt.Println("===================Project Started===================")
}

func WaterDescription(number int) string {
	switch {
	case number <= 5:
		return "Aman."
	case 6 <= number && number <= 8:
		return "Siaga."
	case number > 8:
		return "Bahaya."
	}
	return "Hehe"
}

func WindDescription(number int) string {
	switch {
	case number <= 6:
		return "Aman."
	case 7 <= number && number <= 15:
		return "Siaga."
	case number > 15:
		return "Bahaya."
	}
	return "Wut"
}

func UpdateWeatherData() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// data := map[string]int{
	// 	"water": r1.Intn(20),
	// 	"wind":  r1.Intn(20),
	// }
	water := r1.Intn(20)
	waterDesc := WaterDescription(water)
	waterStr := strconv.Itoa(water)
	wind := r1.Intn(20)
	windDesc := WindDescription(wind)
	windStr := strconv.Itoa(wind)

	// file, _ := json.MarshalIndent(data, "", " ")

	file := fmt.Sprintf(`
	<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="css.bootstrap.css">
			<script src="jquery.js"></script>
		</head>
		<body>
			<h1 center>Dummy Forecast</h1>
			Water : <span id="water">%s</span>
			<br>
			Description : <span id="water-description">%s</span>
			<br>
			Wind : <span id="wind">%s</span>
			<br>
			Description : <span id="wind-description">%s</span>
		</body>
		<script>
			
		</script>
	</html>
	`, waterStr, waterDesc, windStr, windDesc)

	_ = ioutil.WriteFile("forecast.html", []byte(file), 0644)

	fmt.Println("===================New Forecast Data Updated===================")
}
