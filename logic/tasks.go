package logic

import (
	"github.com/bamzi/jobrunner"
	"messag-push/utils"
	"time"
)

func StartTasks() {
	jobrunner.Start()
	jobrunner.Every(1*time.Second, utils.WrapJob("graph_task", GraphTask))
}
