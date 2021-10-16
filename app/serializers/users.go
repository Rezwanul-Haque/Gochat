package serializers

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Sdp struct {
	Sdp string `json:"sdp"`
}
