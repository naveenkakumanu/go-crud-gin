package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// DBConfig
type DBConfig struct {
	DB Database `yaml:"db"`
}

// Database configuration
type Database struct {
	Name       string
	UserName   string
	Credential string
	Host       string
	Driver     string
}

// Reading Database Config from yaml file
func (db *DBConfig) ReadConfig() *DBConfig {
	config, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		log.Printf("Error Reading File", err)
	}
	err = yaml.Unmarshal(config, &db)
	if err != nil {
		log.Printf("Error Unmarshaling File", err)
	}

	return db
}
