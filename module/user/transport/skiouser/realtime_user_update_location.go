package skiouser

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"golang/common"
	"gorm.io/gorm"
	"log"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	GetUser(conn socketio.Conn) common.Requester
}

type UserLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (ul UserLocation) String() string {
	return fmt.Sprintf("Lat: %f, Lng: %f\n", ul.Lat, ul.Lng)
}

func RealtimeUserUploadLocation(appCtx AppContext) func(socketio.Conn, UserLocation) {
	return func(conn socketio.Conn, loc UserLocation) {
		log.Println("RealtimeUserUploadLocation:", loc, ", requester:", appCtx.GetUser(conn).GetUserId()) // who?
	}
}
