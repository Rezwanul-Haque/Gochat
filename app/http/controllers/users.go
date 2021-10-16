package controllers

import (
	"fmt"
	m "gochat/app/http/middlewares"
	"gochat/app/serializers"
	"gochat/app/svc"
	"gochat/app/utils/consts"
	"gochat/infra/errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v2"
)

type users struct {
	uSvc svc.IUsers
}

// NewUsersController will initialize the controllers
func NewUsersController(grp interface{}, uSvc svc.IUsers) {
	uc := &users{
		uSvc: uSvc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/users/signup", uc.Create)
	g.POST("/v1/room", uc.CreateRoom, m.CustomAuth())
	// g.POST("/v1/webrtc/sdp/m/:roomId/c/:userID/p/:peerId/s/:sender", uc.InstantMeeting, m.CustomAuth())
	// g.POST("/v1/webrtc/sdp/m/:meetingId/c/:userID/p/:peerId/s/:isSender", uc.InstantMeeting)
}

func (ctr *users) Create(c echo.Context) error {
	var user serializers.UserReq

	if err := c.Bind(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return c.JSON(restErr.Status, restErr)
	}

	// Password hash is handled by firebase so no need to initiate hash here.
	// hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	// user.Password = string(hashedPass)

	resp, saveErr := ctr.uSvc.CreateUser(user)
	if saveErr != nil {
		return c.JSON(saveErr.Status, saveErr)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (ctr *users) CreateRoom(c echo.Context) error {
	roomID, _ := uuid.NewUUID()

	return c.JSON(http.StatusCreated, map[string]interface{}{"room_id": roomID})
}

// func (ctr *users) InstantMeeting(c echo.Context) error {
// 	isSender, _ := strconv.ParseBool(c.Param("isSender"))
// 	userID := c.Param("userID")
// 	peerID := c.Param("peerId")

// 	var session serializers.Sdp
// 	if err := c.Bind(&session); err != nil {
// 		restErr := errors.NewBadRequestError("invalid json body")
// 		return c.JSON(restErr.Status, restErr)
// 	}
// 	offer := webrtc.SessionDescription{}
// 	methodsutil.Decode(session.Sdp, &offer)

// 	// Create a new RTCPeerConnection
// 	// this is the gist of webrtc, generates and process SDP
// 	peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !isSender {
// 		recieveTrack(peerConnection, peerConnectionMap, peerID)
// 	} else {
// 		createTrack(peerConnection, peerConnectionMap, userID)
// 	}
// 	// Set the SessionDescription of remote peer
// 	peerConnection.SetRemoteDescription(offer)

// 	// Create answer
// 	answer, err := peerConnection.CreateAnswer(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Sets the LocalDescription, and starts our UDP listeners
// 	err = peerConnection.SetLocalDescription(answer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return c.JSON(http.StatusOK, serializers.Sdp{Sdp: methodsutil.Encode(answer)})
// }

// user is the caller of the method
// if user connects before peer: create channel and keep listening till track is added
// if peer connects before user: channel would have been created by peer and track can be added by getting the channel from cache
func recieveTrack(peerConnection *webrtc.PeerConnection,
	peerConnectionMap map[string]chan *webrtc.Track,
	peerID string) {
	if _, ok := peerConnectionMap[peerID]; !ok {
		peerConnectionMap[peerID] = make(chan *webrtc.Track, 1)
	}
	localTrack := <-peerConnectionMap[peerID]
	peerConnection.AddTrack(localTrack)
}

// user is the caller of the method
// if user connects before peer: since user is first, user will create the channel and track and will pass the track to the channel
// if peer connects before user: since peer came already, he created the channel and is listning and waiting for me to create and pass track
func createTrack(peerConnection *webrtc.PeerConnection,
	peerConnectionMap map[string]chan *webrtc.Track,
	currentUserID string) {

	if _, err := peerConnection.AddTransceiver(webrtc.RTPCodecTypeVideo); err != nil {
		log.Fatal(err)
	}

	// Set a handler for when a new remote track starts, this just distributes all our packets
	// to connected peers
	peerConnection.OnTrack(func(remoteTrack *webrtc.Track, receiver *webrtc.RTPReceiver) {
		// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
		// This can be less wasteful by processing incoming RTCP events, then we would emit a NACK/PLI when a viewer requests it
		go func() {
			ticker := time.NewTicker(consts.RtcpPLIInterval)
			for range ticker.C {
				if rtcpSendErr := peerConnection.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: remoteTrack.SSRC()}}); rtcpSendErr != nil {
					fmt.Println(rtcpSendErr)
				}
			}
		}()

		// Create a local track, all our SFU clients will be fed via this track
		// main track of the broadcaster
		localTrack, newTrackErr := peerConnection.NewTrack(remoteTrack.PayloadType(), remoteTrack.SSRC(), "video", "pion")
		if newTrackErr != nil {
			log.Fatal(newTrackErr)
		}

		// the channel that will have the local track that is used by the sender
		// the localTrack needs to be fed to the reciever
		localTrackChan := make(chan *webrtc.Track, 1)
		localTrackChan <- localTrack
		if existingChan, ok := peerConnectionMap[currentUserID]; ok {
			// feed the exsiting track from user with this track
			existingChan <- localTrack
		} else {
			peerConnectionMap[currentUserID] = localTrackChan
		}

		rtpBuf := make([]byte, 1400)
		for { // for publisher only
			i, readErr := remoteTrack.Read(rtpBuf)
			if readErr != nil {
				log.Fatal(readErr)
			}

			// ErrClosedPipe means we don't have any subscribers, this is ok if no peers have connected yet
			if _, err := localTrack.Write(rtpBuf[:i]); err != nil && err != io.ErrClosedPipe {
				log.Fatal(err)
			}
		}
	})
}
