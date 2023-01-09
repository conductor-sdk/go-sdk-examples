package workflow

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"

	"github.com/conductor-sdk/go-sdk-examples/internal/api"
)

var (
	metadataClient = client.MetadataResourceApiService{
		APIClient: api.GetApiClientWithAuthentication(),
	}
	workflowClient = client.WorkflowResourceApiService{
		APIClient: api.GetApiClientWithAuthentication(),
	}
	workflowExecutor = executor.NewWorkflowExecutor(
		api.GetApiClientWithAuthentication(),
	)
)

func ValidateWorkflowDaemon(waitTime time.Duration, outputChannel chan error, workflowId string, expectedOutput map[string]interface{}) {
	time.Sleep(waitTime)
	workflow, _, err := workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		outputChannel <- err
		return
	}
	if !isWorkflowCompleted(&workflow) {
		outputChannel <- fmt.Errorf(
			"workflow status different than expected, workflowId: %s, workflowStatus: %s",
			workflow.WorkflowId, workflow.Status,
		)
		return
	}
	if !reflect.DeepEqual(workflow.Output, expectedOutput) {
		outputChannel <- fmt.Errorf(
			"workflow output is different than expected, workflowId: %s, output: %+v",
			workflow.WorkflowId, workflow.Output,
		)
		return
	}
	outputChannel <- nil
}

func isWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.CompletedWorkflow
}
