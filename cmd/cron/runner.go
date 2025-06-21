package cron

import (
	"github.com/alessandro54/stats/cmd/cron/jobs"
	"github.com/go-co-op/gocron/v2"
)

func StartCronJobs() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	currentSeasonId := "39"

	_, err = jobs.CreatePvpSnapshot(s, currentSeasonId, "2v2")

	_, err = jobs.CreatePvpSnapshot(s, currentSeasonId, "3v3")

	if err != nil {
		panic(err)
	}

	s.Start()
}
