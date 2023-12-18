package main

import (
	"fmt"
	"time"
)

type cook struct {
	name string
	time int
}

const restWorkTime = 10

func startRest() {
	cooks := make(map[string]cook)
	cooks["cook1"] = cook{
		name: "Povar",
		time: 3,
	}
	cooks["cook2"] = cook{
		name: "Povarenok",
		time: 2,
	}

	for _, value := range cooks {
		go func(cook cook) {
			for {
				fmt.Printf("Cook %s started \n", cook.name)
				time.Sleep(time.Duration(cook.time) * time.Second)
				fmt.Printf("Cook %s finished \n", cook.name)
			}
		}(value)
	}

	time.Sleep(time.Duration(restWorkTime) * time.Second)
}

func main() {
	startRest()
}
