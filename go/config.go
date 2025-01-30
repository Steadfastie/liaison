package main

type Config struct {
	MongoSettings struct {
		ConnectionString string `json:"connectionString"`
		Database         string `json:"database"`
	} `json:"MongoSettings"`
	Host struct {
		Port int `json:"port"`
	} `json:"Host"`
}
