package rtcc

type rtcClient struct {
}

var (
	myRtcClient *rtcClient
)

func TokenBuilder() *rtcClient {
	return myRtcClient
}
