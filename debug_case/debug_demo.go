package main

import (
	"flag"
	"fmt"
	"time"
)

var j = flag.Int("j", 0, "")

func main() {
	flag.Parse()
	i := 1
	for {
		fmt.Println("demo ", i, *j)
		i++
		time.Sleep(time.Second)
	}
}
