package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func printStars() {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(5) + 1
	stars := strings.Repeat("*", count)
	fmt.Println(stars)
}
