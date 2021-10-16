package consts

import "time"

const (
	RtcpPLIInterval = time.Second * 3

	// Allows compressing offer/answer to bypass terminal input limits.
	Compress = false
)
