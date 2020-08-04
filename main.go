package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Print("env TZ:", os.Getenv("TZ"))
	log.Print("time.Local:", time.Local)
	log.Println("Local Time:", time.Now())
	log.Println("UTC Time:", time.Now().UTC())
	log.Println("Unix Seconds:", time.Now().Unix())
	log.Println("UnixNano:", time.Now().UnixNano())

	log2 := log.New(os.Stderr, "UTC ", log.LstdFlags|log.LUTC)
	log2.Print("time.Local:", time.Local)
	log2.Println("Local Time:", time.Now())
	log2.Println("UTC Time:", time.Now().UTC())
	log2.Println("Unix Seconds:", time.Now().Unix())
	log2.Println("UnixNano:", time.Now().UnixNano())
}
