package worker

import "github.com/conductor-sdk/conductor-go/sdk/model"

type WorkflowTask struct {
	Name              string `json:"name"`
	TaskReferenceName string `json:"taskReferenceName"`
	Type              string `json:"type,omitempty"`
}

// DynamicForkWorker worker to create a new dynamic task
func DynamicForkWorker(task *model.Task) (result interface{}, err error) {
	taskResult := model.NewTaskResultFromTask(task)
	//List of tasks to be forked
	tasks := []WorkflowTask{
		{
			Name:              "get_user_info",
			TaskReferenceName: "get_user_info",
			Type:              "SIMPLE",
		},
		{
			Name:              "send_email",
			TaskReferenceName: "send_email",
			Type:              "SIMPLE",
		},
		{
			Name:              "send_sms",
			TaskReferenceName: "send_sms",
			Type:              "SIMPLE",
		},
	}
	//Map with a key as task reference name from the above list and value as the input to be given to the task
	inputs := map[string]interface{}{
		"get_user_info": map[string]interface{}{
			"userId": "user_a",
		},
		"send_email": map[string]interface{}{
			"email":   "user@example.com",
			"subject": "Hello",
		},
		"send_sms": map[string]interface{}{
			"phoneNumber": "555-555-5555",
			"text":        "Hi",
		},
	}

	taskResult.OutputData = map[string]interface{}{
		"forkedTasks":       tasks,
		"forkedTasksInputs": inputs,
	}
	taskResult.Status = model.CompletedTask
	err = nil
	return taskResult, err
}
