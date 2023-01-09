package worker

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func GetUserInfo(task *model.Task) (interface{}, error) {
	userId, err := GetValueFromTaskInputData(task, "userId")
	if err != nil {
		return nil, err
	}
	userInfo := NewUserInfo("User X", *userId)
	userInfo.Email = fmt.Sprintf("%s@example.com", *userId)
	userInfo.PhoneNumber = "555-555-5555"
	return userInfo, nil
}

func SendEmail(task *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(task)
	email, err := GetValueFromTaskInputData(task, "email")
	if err != nil {
		return nil, err
	}
	taskResult.Logs = []model.TaskExecLog{
		{
			Log: fmt.Sprintf("sent email to %s", *email),
		},
	}
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}

func SendSms(task *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(task)
	phoneNumber, err := GetValueFromTaskInputData(task, "phoneNumber")
	if err != nil {
		return nil, err
	}
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log: fmt.Sprintf("sent sms to %s", *phoneNumber),
		},
	)
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}
