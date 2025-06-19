package cron

import (
	"context"
	"fmt"
	"github.com/alessandro54/stats/internal/gameinfo/domain/port"
	"github.com/go-co-op/gocron/v2"
)

func StartCronJobs(leaderboardSnapshot port.LeaderboardSnapshotService) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	j, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(gocron.NewAtTime(12, 0, 0)),
		),
		gocron.NewTask(
			func() {
				ctx := context.Background()
				err := leaderboardSnapshot.FetchFromBlizzardAndSave(ctx, "33", "3v3")
				if err != nil {
					fmt.Println("Error fetching snapshots:", err)
					return
				}
			},
		),
	)
	fmt.Println(j.ID())
	if err != nil {
		panic(err)
	}

	s.Start()
}
