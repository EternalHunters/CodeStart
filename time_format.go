package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02T15:04:05+08:00"))
}
