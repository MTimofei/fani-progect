package main

import (
	"time"

	"github.com/fani_progect/incert/get/addr"
)

func main() {

	for {
		u := addr.New("https://www.youtube.com")
		u.Start()
		time.Sleep(2 * time.Second)
	}

}

//test
