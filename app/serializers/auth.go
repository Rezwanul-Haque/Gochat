package serializers

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
