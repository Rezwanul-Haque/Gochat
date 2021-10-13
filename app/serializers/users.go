package serializers

type UserReq struct {
	Email       string `json:"email"`
	Phone       string `json:"phone,omitempty"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name,omitempty"`
	ProfilePic  string `json:"profile_pic,omitempty"`
}

type UserResp map[string]interface{}

type LoggedInUser struct {
	ID          int    `json:"user_id"`
	AccessUuid  string `json:"access_uuid"`
	RefreshUuid string `json:"refresh_uuid"`
}
