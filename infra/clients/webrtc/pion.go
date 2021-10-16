package webrtc

import "github.com/pion/webrtc/v2"

type Event struct {
	PeerConnectionMap map[string]chan *webrtc.Track
	Me                *webrtc.MediaEngine
	Api               *webrtc.API
	Config            *webrtc.Configuration
}

func Init() {
	var evt Event

	// sender to channel of track
	evt.PeerConnectionMap = make(map[string]chan *webrtc.Track)

	evt.Me = &webrtc.MediaEngine{}

	// Setup the codecs you want to use.
	// Only support VP8(video compression), this makes our proxying code simpler
	evt.Me.RegisterCodec(webrtc.NewRTPVP8Codec(webrtc.DefaultPayloadTypeVP8, 90000))

	evt.Api = webrtc.NewAPI(webrtc.WithMediaEngine(*evt.Me))

	evt.Config = &webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"}, // public STUN server
			},
		},
	}
}
