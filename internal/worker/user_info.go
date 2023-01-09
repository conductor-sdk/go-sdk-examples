package worker

type UserInfo struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

func NewUserInfo(name string, id string) *UserInfo {
	return &UserInfo{
		Name: name,
		Id:   id,
	}
}
