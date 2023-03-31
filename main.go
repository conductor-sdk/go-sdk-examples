package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"

	apiUtil "github.com/conductor-sdk/go-sdk-examples/internal/api"
	workerUtil "github.com/conductor-sdk/go-sdk-examples/internal/worker"
	workflowUtil "github.com/conductor-sdk/go-sdk-examples/internal/workflow"
)

func main() {
	setupLogSettings()
	workerUtil.StartWorkers()
	wf := workflowUtil.CreateAndRegisterWorkflow()
	log.Info("\n\nStarting a workflow synchronously")
	startWorkflowSync(wf)

	log.Info("\n\nStarting a workflow asynchronously")
	startWorkflowAsync(wf)

	log.Info("\n\nStarting a workflow with Dynamic Fork/Join")
	dynWorkflow := workflowUtil.CreateAndRegisterDynamicForkWorkflow()
	startWorkflowSync(dynWorkflow)

	workflowUtil.StartWorkflowWithRequest()
}

func setupLogSettings() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Info("Finished setting up log settings")
}

func startWorkflowSync(wf *workflow.ConductorWorkflow) {
	workflowInput := workflowUtil.NewWorkflowInput("userA")
	workflowRun, err := wf.ExecuteWorkflowWithInput(workflowInput, "")
	if err != nil {
		panic("failed to execute sync workflow, err: " + fmt.Sprint(err))
	}
	log.Info()
	log.Info("=======================================================================================")
	log.Info("Workflow Execution Completed")
	log.Info("Workflow Id: " + workflowRun.WorkflowId)
	log.Info("Workflow Status: " + workflowRun.Status)
	log.Info("Workflow Output: " + fmt.Sprint(workflowRun.Output))
	log.Info("Workflow Execution Flow UI: " + apiUtil.GetWorkflowExecutionURL(workflowRun.WorkflowId))
	log.Info("=======================================================================================")
	if workflowRun.Status != string(model.CompletedWorkflow) {
		panic("workflow not completed")
	}
}

func startWorkflowAsync(wf *workflow.ConductorWorkflow) {
	workflowInput := workflowUtil.NewWorkflowInput("userA")
	workflowId, err := wf.StartWorkflowWithInput(workflowInput)
	if err != nil {
		panic("failed to start async workflow, err: " + fmt.Sprint(err))
	}
	time.Sleep(5 * time.Second)
	log.Info()
	log.Info("=======================================================================================")
	log.Info("Workflow Execution Completed")
	log.Info("Workflow Id: " + workflowId)
	log.Info("Workflow Execution Flow UI: " + apiUtil.GetWorkflowExecutionURL(workflowId))
	log.Info("=======================================================================================")
}
