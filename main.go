package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	run()
}

func run(){
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 2
	fmt.Println(rand.Intn(max-min+1) + min)
}