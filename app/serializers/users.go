package serializers

import (
	"time"
)

type UserReq struct {
	UserName   string  `json:"user_name,omitempty"`
	FirstName  string  `json:"first_name,omitempty"`
	LastName   string  `json:"last_name,omitempty"`
	Email      string  `json:"email,omitempty"`
	Password   *string `json:"password,omitempty"`
	ProfilePic *string `json:"profile_pic,omitempty"`
	Phone      string  `json:"phone,omitempty"`
	CompanyID  uint    `json:"company_id"`
}

type LoggedInUser struct {
	ID          int    `json:"user_id"`
	AccessUuid  string `json:"access_uuid"`
	RefreshUuid string `json:"refresh_uuid"`
}

type UserResp struct {
	ID          int        `json:"id"`
	UserName    string     `json:"user_name"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	CompanyID   uint       `json:"company_id"`
	Email       string     `json:"email"`
	Phone       *string    `json:"phone"`
	ProfilePic  *string    `json:"profile_pic"`
	AppKey      string     `json:"app_key,omitempty"`
	RoleID      uint       `json:"role_id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	LastLoginAt *time.Time `json:"last_login_at"`
	FirstLogin  bool       `json:"first_login"`
}
