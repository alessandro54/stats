package jobs

import (
	"context"
	"fmt"
	"github.com/alessandro54/stats/internal/gameinfo/persistence/repositories"
	"github.com/alessandro54/stats/internal/gameinfo/services"
	"github.com/go-co-op/gocron/v2"
)

func CreatePvpSnapshot(scheduler gocron.Scheduler, pvpSeasonId string, pvpBracket string) (gocron.Job, error) {
	repo := repositories.NewLeaderboardSnapshotRepository()
	service := services.NewSnapshotService(repo)

	job, err := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0)),
		),
		gocron.NewTask(
			func() {
				ctx := context.Background()
				err := service.FetchFromBlizzardAndSave(ctx, pvpSeasonId, pvpBracket)
				if err != nil {
					fmt.Println("Error fetching snapshots:", err)
					return
				}
			},
		),
	)
	// print job id

	fmt.Println(job.ID())

	return job, err
}
