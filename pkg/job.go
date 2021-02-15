package pkg

import (
	"github.com/Kuzmrom7/ping-go/config"
	"github.com/jasonlvhit/gocron"
)

func Start(jobFun interface{}, cfg *config.Configurations) {
	s := gocron.NewScheduler()
	_ = s.Every(cfg.Seconds).Seconds().Do(jobFun, cfg)

	<-s.Start()

	defer gocron.Clear()
}
