package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/conductor-sdk/conductor-go/sdk/workflow"

	"github.com/conductor-sdk/go-sdk-examples/internal/worker"
	workflowUtil "github.com/conductor-sdk/go-sdk-examples/internal/workflow"
)

func main() {
	setupLogSettings()
	worker.StartWorkers()
	wf := workflowUtil.CreateAndRegisterWorkflow()
	startWorkflowSync(wf)
	startWorkflowAsync(wf)
}

func setupLogSettings() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.Info("Finished setting up log settings")
}

func startWorkflowSync(wf *workflow.ConductorWorkflow) {
	log.Info()
	log.Info("=======================================================================================")
	log.Info("Workflow Execution Completed")
	// log.Info("Workflow Id: " + workflowRun.WorkflowId)
	// log.Info("Workflow Status: " + workflowRun.Status.ToString())
	// log.Info("Workflow Output: " + workflowRun.Output.ToString())
	// log.Info("Workflow Execution Flow UI: {Examples.Api.ApiUtil.GetWorkflowExecutionURL(workflowRun.WorkflowId)}")
	log.Info("=======================================================================================")
}

func startWorkflowAsync(wf *workflow.ConductorWorkflow) {
	time.Sleep(5 * time.Second)
	log.Info()
	log.Info("=======================================================================================")
	log.Info("Workflow Execution Completed")
	// log.Info("Workflow Id: {workflowId}")
	// log.Info("Workflow Execution Flow UI: {Examples.Api.ApiUtil.GetWorkflowExecutionURL(workflowId)}")
	log.Info("=======================================================================================")
}
