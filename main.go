package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)
type apiConfigData struct{
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey`
}
type weatherData struct{
	Name string `json:"name`
	Main struct{
		kelvin float64 `json:"temp`
	}`json:"main`
}
func loadApiKey(filename string)(apiConfigData, error){
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil{
		return apiConfigData{}, err
	}

	return c, nil
}