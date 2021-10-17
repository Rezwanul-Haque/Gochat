package serializers

import (
	"fmt"
	"gochat/app/utils/methodsutil"
	"gochat/infra/logger"
	"sync"

	"github.com/gorilla/websocket"
)

type RoomResp struct {
	RoomID string `json:"room_id"`
}

// Participant describes a single entity in the hashmap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type BroadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

// RoomMap is the main hashmap [roomID string] -> [[]Participant]
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// Init initialises the RoomMap struct
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// Get will return the array of participants in the room
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

// CreateRoom generate a unique room ID and return it -> insert it in the hashmap
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	roomID := methodsutil.RandomText(8)

	r.Map[roomID] = []Participant{}

	return roomID
}

// InsertIntoRoom will create a participant and add it in the hashmap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	logger.Info(fmt.Sprintf("inserting into room with roomid: %v", roomID))
	r.Map[roomID] = append(r.Map[roomID], p)
}

// DeleteRoom deletes the room with the roomID
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
