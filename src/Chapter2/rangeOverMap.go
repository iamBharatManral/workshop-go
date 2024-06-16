package main

import "fmt"

func rangeOverMap() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.0",
	}
	for key, value := range config {
		fmt.Println(key, " = ", value)
	}
}
