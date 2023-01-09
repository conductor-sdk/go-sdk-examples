package workflow

const (
	Email string = "EMAIL"
	Sms   string = "SMS"
)

type WorkflowInput struct {
	UserId           string `json:"userId"`
	NotificationPref string `json:"notificationPref"`
}

func NewWorkflowInput(userId string) *WorkflowInput {
	return &WorkflowInput{
		UserId:           userId,
		NotificationPref: Sms,
	}
}
