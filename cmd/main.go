package main

import (
	"fmt"
	"strconv"
)

func main() {
	appConfig := SetupConfig()
	fmt.Println("Configuration Loaded")
	property, err := appConfig.GetValue("USERNAME", IntegerParser())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Print something %s\n with Type %t", property, property)
}

func IntegerParser() PropertyParser {
	return func(s interface{}) (interface{}, error) {
		str := s.(string)
		val, err := strconv.Atoi(str)
		return val, err
	}
}
