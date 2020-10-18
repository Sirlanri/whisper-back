package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	where := make([]string, 0, 3)
	where = append(where, "1", "2", "3")
	rand.Seed(time.Now().Unix())
	result := where[rand.Intn(4)]
	fmt.Println("今晚去", result)
	return
}
