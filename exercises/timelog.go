package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// const sec1 = 5
// const sec2 = 10
// const sec3 = 15

func main() {

	goPrintLog()
	for {
		time.Sleep(time.Second)
	}

}

func goPrintLog() {
	c := cron.New()
	c.AddFunc("@every 5s", func() {
		log.Printf("5sec")
	})
	c.AddFunc("@every 10s", func() {
		log.Printf("10sec")
	})
	c.AddFunc("@every 15s", func() {
		log.Printf("15sec")
	})
	c.Start()

}

// func goPrintLog() {

// 	go func() {
// 		for {
// 			time.Sleep(sec1 * time.Second)
// 			log.Println("5sec log")

// 		}

// 	}()

// 	go func() {
// 		for {
// 			time.Sleep(sec2 * time.Second)
// 			log.Println("10sec log")

// 		}

// 	}()

// 	func() {
// 		for {
// 			time.Sleep(sec3 * time.Second)
// 			log.Println("15sec log")

// 		}

// 	}()

// }
