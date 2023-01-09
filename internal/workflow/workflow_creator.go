package workflow

import (
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
)

func CreateAndRegisterWorkflow() *workflow.ConductorWorkflow {
	wf := createComplexWorkflow()
	workflowExecutor.RegisterWorkflow(
		true,
		wf.ToWorkflowDef(),
	)
	return wf
}

func createComplexWorkflow() *workflow.ConductorWorkflow {
	return workflow.NewConductorWorkflow(workflowExecutor).
		Name("user_notification").
		Version(1).
		InputParameters("userId", "notificationPref").
		Add(createGetUserDetailsTask()).
		Add(createEmailOrSMSTask())
}

func createGetUserDetailsTask() workflow.TaskInterface {
	return workflow.NewSimpleTask("get_user_info", "get_user_info").
		Input("userId", "${workflow.input.userId}")
}

func createEmailOrSMSTask() workflow.TaskInterface {
	return workflow.NewSwitchTask("emailorsms", "${workflow.input.notificationPref}").
		SwitchCase(Email, createSendEmailTask()).
		SwitchCase(Sms, createSendSMSTask())
}

func createSendEmailTask() workflow.TaskInterface {
	return workflow.NewSimpleTask("send_email", "send_email").
		Input("email", "${get_user_info.output.email}")
}

func createSendSMSTask() workflow.TaskInterface {
	return workflow.NewSimpleTask("send_sms", "send_sms").
		Input("phoneNumber", "${get_user_info.output.phoneNumber}")
}
