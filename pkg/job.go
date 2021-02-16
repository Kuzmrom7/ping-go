package pkg

import (
	"github.com/Kuzmrom7/ping-go/config"
	"github.com/jasonlvhit/gocron"
	"log"
)

func RunJob(jobFun interface{}, cfg *config.Configurations) {
	s := gocron.NewScheduler()
	err := s.Every(cfg.Timeout).Minutes().Do(jobFun, cfg)
	if err != nil {
		log.Println("Job creating error")

		return
	}

	<-s.Start()

	defer gocron.Clear()
}
