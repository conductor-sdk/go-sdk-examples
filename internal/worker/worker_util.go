package worker

import (
	"fmt"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	log "github.com/sirupsen/logrus"

	"github.com/conductor-sdk/go-sdk-examples/internal/api"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(
	api.GetApiClientWithAuthentication(),
)

const (
	batchSize    = 5
	pollInterval = 100 * time.Millisecond
)

func StartWorkers() {
	startWorker("get_user_info", GetUserInfo)
	startWorker("send_email", SendEmail)
	startWorker("send_sms", SendSms)
	log.Info("started all workers")
}

func GetValueFromTaskInputData(t *model.Task, key string) (value *string, err error) {
	rawValue, ok := t.InputData[key]
	if !ok {
		return nil, fmt.Errorf("key (%s) is not set", key)
	}
	*value, ok = rawValue.(string)
	if !ok {
		return nil, fmt.Errorf("key (%s) value must be of string type", key)
	}
	return value, nil
}

func startWorker(taskName string, executeFunction model.ExecuteTaskFunction) {
	taskRunner.StartWorker(
		taskName,
		executeFunction,
		batchSize,
		pollInterval,
	)
}
