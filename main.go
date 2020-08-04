package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log2 := log.New(os.Stderr, "UTC ", log.LstdFlags|log.LUTC)

	c := time.NewTicker(10 * time.Second).C
	for ; ; <-c {
		log.Print("env TZ:", os.Getenv("TZ"))
		log.Print("time.Local:", time.Local)
		log.Print("Local Time:", time.Now())
		log.Print("UTC Time:", time.Now().UTC())
		log.Print("Unix Seconds:", time.Now().Unix())
		log.Print("UnixNano:", time.Now().UnixNano())

		log2.Print("time.Local:", time.Local)
		log2.Print("Local Time:", time.Now())
		log2.Print("UTC Time:", time.Now().UTC())
		log2.Print("Unix Seconds:", time.Now().Unix())
		log2.Print("UnixNano:", time.Now().UnixNano())
	}
}
