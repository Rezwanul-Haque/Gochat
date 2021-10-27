package controllers

import (
	m "gochat/app/http/middlewares"
	"gochat/app/serializers"
	"gochat/infra/logger"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type rooms struct{}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// AllRooms is the global hashmap for the server
var AllRooms serializers.RoomMap

// NewRoomsController will initialize the controllers
func NewRoomsController(grp interface{}) {
	rc := &rooms{}

	g := grp.(*echo.Group)

	AllRooms.Init()

	g.GET("/v1/room", rc.CreateRoom, m.CustomAuth())
	// g.GET("/v1/join", rc.JoinRoom)
	// g.GET("/v1/room", rc.CreateRoom) // testing purposes only
	// g.GET("/v1/join", rc.JoinRoom) // testing purposes only
}

// CreateRoom Create a Room and return roomID
func (ctr *rooms) CreateRoom(c echo.Context) error {
	roomID := AllRooms.CreateRoom()

	logger.InfoAsJson("all rooms map", AllRooms.Map)
	return c.JSON(http.StatusOK, serializers.RoomResp{RoomID: roomID})
}

// JoinRoom will join the client in a particular room
// func (ctr *rooms) JoinRoom(c echo.Context) error {
// 	roomID := c.QueryParam("roomID")

// 	if roomID == "" {
// 		logger.Info("roomID missing in url parameters")
// 		restErr := errors.NewBadRequestError("room id cannot be empty")
// 		return c.JSON(restErr.Status, restErr)
// 	}

// 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		logger.Error("web socket upgrade error", err)
// 		return err
// 	}
// 	defer ws.Close()

// 	AllRooms.InsertIntoRoom(roomID, false, ws)

// 	go broadcaster()

// 	for {
// 		var msg serializers.BroadcastMsg

// 		err := ws.ReadJSON(&msg.Message)
// 		if err != nil {
// 			logger.Error("json read error: ", err)
// 			restErr := errors.NewInternalServerError(errors.ErrSomethingWentWrong)
// 			return c.JSON(restErr.Status, restErr)
// 		}

// 		msg.Client = ws
// 		msg.RoomID = roomID

// 		logger.InfoAsJson("web socket received message", msg)

// 		broadcast <- msg
// 	}
// }

// var broadcast = make(chan serializers.BroadcastMsg)

// func broadcaster() {
// 	for {
// 		msg := <-broadcast

// 		for _, client := range AllRooms.Map[msg.RoomID] {
// 			if client.Conn != msg.Client {
// 				err := client.Conn.WriteJSON(msg.Message)

// 				if err != nil {
// 					logger.Error("writing on json", err)
// 					client.Conn.Close()
// 					return
// 				}
// 			}
// 		}
// 	}
// }
