package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, World",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا",
	"Привет, мир",
}

func randomMessgae() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList) {
		return "", errors.New("out of range" + strconv.Itoa(index))
	}
	return helloList[index], nil
}
