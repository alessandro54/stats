package jobs

import (
	"context"
	"fmt"
	"github.com/alessandro54/stats/internal/dataextraction/service"
	"github.com/alessandro54/stats/internal/infra/db"
	"github.com/alessandro54/stats/internal/playervsplayer/persistence/repositories"
	"github.com/go-co-op/gocron/v2"
)

func CreatePvpSnapshot(scheduler gocron.Scheduler, pvpSeasonId string, pvpBracket string, region string) (gocron.Job, error) {
	repo := repositories.NewSnapshotRepository(db.ProvideDB())
	svc := service.NewSnapshotService(repo)

	job, err := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0)),
		),
		gocron.NewTask(
			func() {
				ctx := context.Background()
				err := svc.FetchFromBlizzardAndSave(ctx, pvpSeasonId, pvpBracket, region)
				if err != nil {
					fmt.Println("Error fetching snapshots:", err)
					return
				}
			},
		),
	)

	fmt.Println(job.ID())

	return job, err
}
