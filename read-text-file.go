package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type City struct {
	Name string `json:"name"`
}
type Country struct {
	Name string `json:"name"`
}

func cities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// open the file using Open() function from os library
	file, err := os.Open("world-cities.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	cities := []City{}
	for scanner.Scan() {
		cityJson := City{
			Name: scanner.Text(),
		}
		cities = append(cities, cityJson)
	}

	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(cities)
	file.Close()
}

func countries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// open the file using Open() function from os library
	file, err := os.Open("countries.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read the file line by line using a scanner
	scanner := bufio.NewScanner(file)
	countries := []Country{}
	for scanner.Scan() {
		countryJson := Country{
			Name: scanner.Text(),
		}
		countries = append(countries, countryJson)
	}

	// check for the error that occurred during the scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(countries)
	file.Close()
}

func main() {
	http.HandleFunc("/cities", cities)
	http.HandleFunc("/countries", countries)
	http.ListenAndServe(":8080", nil)
}
