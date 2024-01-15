package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	persentTime := time.Now()
	fmt.Println(persentTime.Format("01-02-2006"))
}
