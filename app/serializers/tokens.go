package serializers

type TokenReq struct {
	ChannelName string `json:"channel_name"`
	TokenType   string `json:"token_type"`
	UID         string `json:"uid"`
	Role        string `json:"role"`
	ExpireIn    uint32 `json:"expire_in,omitempty"`
}

type TokenResp struct {
	RtcToken string `json:"rtc_token"`
}
