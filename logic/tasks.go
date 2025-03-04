package logic

import (
	"github.com/bamzi/jobrunner"
	"messag-push/utils"
	"time"
)

func StartTasks() {
	jobrunner.Start()
	jobrunner.Every(10*time.Second, utils.WrapJob("remote_config_task", FetchAlertConfig))
	jobrunner.Every(1*time.Second, utils.WrapJob("graph_task", GraphTask))
	jobrunner.Every(5*time.Second, utils.WrapJob("currentCap_task", CurrentCapTask))
}
