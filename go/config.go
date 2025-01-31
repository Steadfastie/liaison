package main

type Config struct {
	MongoSettings struct {
		ConnectionString string `json:"ConnectionString"`
		Database         string `json:"Database"`
	} `json:"MongoSettings"`
	Host struct {
		Port int `json:"Port"`
	} `json:"Host"`
}
