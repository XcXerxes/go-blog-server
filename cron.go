package main

import (
	"log"

	"github.com/XcXerxes/go-blog-server/models"
	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("starting...")

	c := cron.New()
	c.AddFunc("******", func() {
		log.Println("Run models.CleanAllTag...")
		models.ClearAllTag()
	})
	c.AddFunc("******", func() {
		log.Panicln("Run models.CleanAllArticle...")
		models.ClearAllArticle()
	})
	c.Start()

	// t1 := time.NewTimer(time.Second * 10)
	// for {
	// 	select {
	// 	case <-t1.C:
	// 		t1.Reset(time.Second * 10)
	// 	}
	// }
}
