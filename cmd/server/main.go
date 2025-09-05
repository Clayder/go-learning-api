package main

import "github.com/Clayder/go-learning-api/configs"

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	println(config.DBDriver)
}	