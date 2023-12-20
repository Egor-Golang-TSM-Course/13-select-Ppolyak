package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

type Restaurant struct {
	name     string
	quantity int
}

func startSale() {
	rest := make(map[string]*Restaurant)
	rest["r1"] = &Restaurant{name: "china"}
	rest["r2"] = &Restaurant{name: "chiza"}
	rest["r3"] = &Restaurant{name: "chizes"}

	res := make(chan Restaurant)
	go centralOf(res)
	for _, val := range rest {
		go func(restaurant *Restaurant) {
			for {
				rand := rand2.Int()
				time.Sleep(2 * time.Second)
				restaurant.quantity += rand
				res <- *restaurant
			}
		}(val)
	}
	time.Sleep(5 * time.Second)
	close(res)
}

func centralOf(res chan Restaurant) {
	for r := range res {
		fmt.Printf("Sales %+v \n", r)
	}
}

func main() {
	startSale()
}
