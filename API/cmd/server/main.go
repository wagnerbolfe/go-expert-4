package main

import "github.com/wagnerbolfe/goexpert/API/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
