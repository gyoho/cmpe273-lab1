package main

import (
	"fmt"
	"time"
)

func sleep(sec int64) {
	<- time.After(time.Duration(sec)*time.Second)
}

func main() {
	startTime := time.Now()
	sleep(10)
	elapsed := time.Since(startTime)
	
	fmt.Println(elapsed)
}