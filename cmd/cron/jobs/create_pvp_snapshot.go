package jobs

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
)

func CreatePvpSnapshot(scheduler gocron.Scheduler, pvpSeasonId string, pvpBracket string, region string) (gocron.Job, error) {

	job, err := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0)),
		),
		gocron.NewTask(
			func() {
				//err := svc.FetchFromBlizzardAndSave(ctx, pvpSeasonId, pvpBracket, region)
				//if err != nil {
				//	fmt.Println("Error fetching snapshots:", err)
				//	return
				//}
				println("Starting Pvp Snapshot")
			},
		),
	)

	fmt.Println(job.ID())

	return job, err
}
