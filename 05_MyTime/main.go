package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// We have to set 02-01-2006 as date to specify the date format :)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	fmt.Println()

	createdDate := time.Date(2020, time.April, 15, 14, 56, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))

}
