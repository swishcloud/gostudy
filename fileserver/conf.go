package main

import (
	"encoding/json"
	"os"
)
type Config struct{
	FileLocation string
	Host string
}
var config Config
func ReadConfig() (Config) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	dec := json.NewDecoder(file)
	var v  map[string]interface{}
	dec.Decode(&v)
	return Config{FileLocation:v["FileLocation"].(string), Host:v["Host"].(string)}
}
